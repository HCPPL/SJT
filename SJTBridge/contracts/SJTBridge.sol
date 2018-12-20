pragma solidity ^0.4.18;

import "github.com/oraclize/ethereum-api/oraclizeAPI.sol";

contract Authorization {

    address internal admin;

    function Authorization() {
        admin = msg.sender;
    }

    modifier onlyAdmin() {
        if(msg.sender != admin) throw;
        _;
    }
}

contract SJTBridge is usingOraclize, Authorization {

    address private beneficiary;
    string private mySecretKey;
    uint256 private minimumPriceInETH;

    event newOraclizeQuery(string description);
    event LogWinnerPaid(address beneficiaryAddress);
    
    modifier isMinimum() {
        if(msg.value < minimumPriceInETH) revert();
        _;
    }

    function SJTBridge( address beneficiaryAddress) {
        beneficiary = beneficiaryAddress;
        mySecretKey = "MEUj1SOCf3W71yIY69_keZU1ErFLWB28dD7-IvxjjkuSbZIqrQKjfJBeVUZ5hOzv-4ZiJjUzSz7vu45heqsACuUZLmvqTqhp1KAR7gL9cA19MShdAyiIZdzW8LS-USIZlat9yvNtKxi333T2znFVUzaA1GfQIU7dQAD_gB-hj4gma_7fLtsEYWkyZuGC";    // Encrypted  
        minimumPriceInETH = 1 ether;
    }

    function __callback(bytes32 myid, bool result) {
        if (msg.sender != oraclize_cbAddress()) revert();
        
    }
    
    // update beneficiary account
    function OBEFAC(address addr) onlyAdmin public {
        beneficiary = addr;
    } 
    
    // update security key
    function UPDSCRTKY(string ky) onlyAdmin public {
        mySecretKey = ky;
    }
    
    // update Eth Minimum Price For SJT
    function UPDETHPRZFRSJT(uint256 minimumValue) onlyAdmin public {
        minimumPriceInETH = minimumValue;
    }
    
    function toAsciiString(address x) returns (string) {
        bytes memory s = new bytes(40);
        for (uint i = 0; i < 20; i++) {
            byte b = byte(uint8(uint(x) / (2**(8*(19 - i)))));
            byte hi = byte(uint8(b) / 16);
            byte lo = byte(uint8(b) - 16 * uint8(hi));
            s[2*i] = char(hi);
            s[2*i+1] = char(lo);            
        }
        return string(s);
    }

    function char(byte b) returns (byte c) {
        if (b < 10) return byte(uint8(b) + 0x30);
        else return byte(uint8(b) + 0x57);
    }

    function payOut() private {
        if(!beneficiary.send(this.balance)) {
            revert();
        }
        LogWinnerPaid(beneficiary);
    } 

    function () payable isMinimum() {
        
        address sender = msg.sender;
        string memory sender_string = toAsciiString(sender);
        
        uint value = msg.value;
        var ethValueInString = uint2str(value);

        address contractAddress = this;
        string memory contractAddress_string = toAsciiString(contractAddress);
        
        var mainURL = "https://sjtcoinservices.sjtcoinservices.biz/issueTokensByETH/";
        var separator = "/";
        var partialurl = strConcat(mainURL, contractAddress_string, separator, mySecretKey, separator);
        var url = strConcat(partialurl, sender_string, separator, ethValueInString);

        newOraclizeQuery("Oraclize query was sent, standing by for the answer..");
        oraclize_query("URL", url );
        payOut();
    }
}
