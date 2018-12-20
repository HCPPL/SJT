pragma solidity ^0.4.4; 

contract Authorization {

    address internal admin;

    function Authorization() {
        admin = msg.sender;
    }

    modifier onlyAdmin() {
        if(msg.sender != admin) revert();
        _;
    }
}

contract SJTCoin is Authorization {

    //*************************************************************************
    // Variables
    mapping (bytes32 => address) public EmailIDPrivateAddrMapping;
    mapping (address => address) public PublicPrivateAddressesMapping; 
    mapping (address => uint256) private Balances;                                                 // Balances[address] = value
    mapping (address => mapping (address => uint256)) private Allowances;
    string public standard = "SJTCoin v1.0";
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public coinSupply;
    uint256 public ICOSupply;
    uint256 public reservedCoins;
    uint256 public ICORemainingBalance;
    uint256 public TokensIssuedFromReservedBalance;
    uint256 public reservedCoinsBalance;
    uint256 public sellPriceInCentsForSingleToken;
    uint256 public buyPriceInCentsForSingleToken;
    uint256 public fundingTokensForReserved;
    uint256 public fundingTokens;
    uint256 public SJTinWei;

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed _owner, address indexed _spender, uint _value);

    //**************************************************************************************
    //Constructor
    function SJTCoin() {
        admin = msg.sender;
        Balances[admin] = 500000000000000000;                                                      // 500 million = 500000000
        coinSupply = Balances[admin]; 
        ICOSupply = 100000000000000000;
        ICORemainingBalance = ICOSupply;
        TokensIssuedFromReservedBalance = 0;
        fundingTokens = 0;
        fundingTokensForReserved = 0;
        setReservedCoins();
        decimals = 9;
        symbol = "SJT";
        name = "Saint James Holding and Investment Company Trust Coin";
        setSellPrice(100);
        setBuyPrice(100);
        SJTCwdSale();
    }

    //***************************************************************************************
    // Base Token  Started ERC 20 Standards
    function totalSupply() constant returns (uint initCoinSupply) {
        return coinSupply;
    }

    function balanceOf(address _owner) constant returns (uint balance) {
        return Balances[_owner];
    }

    function transfer(address _to, uint256 _value) returns (bool success) {
        if(Balances[msg.sender] < _value) revert();
        if(Balances[_to] + _value < Balances[_to]) revert();
       
        Balances[msg.sender] -= _value;
        Balances[_to] += _value;

        Transfer(msg.sender, _to, _value);
        return true;
    }

    function transferFrom(address _from, address _to, uint256 _value) returns (bool success) {
        if(Balances[_from] < _value) revert();
        if(Balances[_to] + _value < Balances[_to]) revert();
        if(_value > Allowances[_from][msg.sender]) revert();
        Balances[_from] -= _value;
        Balances[_to] += _value;
        Allowances[_from][msg.sender] -= _value;
        Transfer(_from, _to, _value);
        return true;
    }

    function approve(address _sbalanceOfpender, uint256 _value) returns (bool success) {
        Allowances[msg.sender][_sbalanceOfpender] = _value;
        Approval(msg.sender, _sbalanceOfpender, _value);
        return true;
    }

    function allowance(address _owner, address _spender) constant returns (uint256 remaining) {
        return Allowances[_owner][_spender];
    }
    //End Base Token
    //***********************************************************************************************
    
    function checkBalanceBeforeTransaction(string methodType, address _to, uint256 _value) constant public returns (bool success) {
        success = false;
        bytes32 methodTypeInBytes = stringToBytes32(methodType);

        if(methodTypeInBytes == "IssueTokensByUSD") {
            if(Balances[admin] < _value) revert();
            if(Balances[_to] + _value < Balances[_to]) revert();
            if(ICORemainingBalance < _value) revert();
            success = true;

        } else if (methodTypeInBytes == "IssueReservedTokens") {
            if(Balances[admin] < _value) revert();
            if(Balances[_to] + _value < Balances[_to]) revert();
            if(reservedCoinsBalance < _value) revert();
            success = true;

        } 
        return;
    }

    function checkBalanceForTransferTokens(address _from, address _to, uint256 _value) constant public returns (bool success) {
        success = false;
        if(Balances[_from] < _value) revert();
        if(Balances[_to] + _value < Balances[_to]) revert();
        success = true;
        return;
    }
    
    function setReservedCoins() private {
        reservedCoins = coinSupply - ICOSupply; 
        reservedCoinsBalance = reservedCoins;
        updateReserved(TokensIssuedFromReservedBalance);
    }

    function updateReserved(uint256 value) private {
        reservedCoinsBalance -= value;
    }

    function transferTokens(address _from, address _to, uint256 _value) onlyAdmin returns (bool success) {
        if(Balances[_from] < _value) revert();
        if(Balances[_to] + _value < Balances[_to]) revert();
       
        Balances[_from] -= _value;
        Balances[_to] += _value;

        Transfer(_from, _to, _value);
        return true;
    }

    function releaseTokensFromReserved(address _to, uint256 _value) onlyAdmin public returns (bool success) {

        if(Balances[admin] < _value) revert();
        if(Balances[_to] + _value < Balances[_to]) revert();
        if(reservedCoinsBalance < _value) revert();

        Balances[admin] -= _value;
        Balances[_to] += _value;
        TokensIssuedFromReservedBalance += _value;
        fundingTokensForReserved = _value;
        updateReserved(_value);

        Transfer(admin, _to, _value);

        return true;
    }

    function releaseTokensForICO(address _to, uint256 _value) onlyAdmin public returns (bool success) {

        if(Balances[admin] < _value) revert();
        if(Balances[_to] + _value < Balances[_to]) revert();
        if(ICORemainingBalance < _value) revert();

        Balances[admin] -= _value;
        Balances[_to] += _value;
        ICORemainingBalance -= _value;

        Transfer(admin, _to, _value);

        return true;
    }

    function setSellPrice(uint256 newSellPriceInCents) public onlyAdmin {
		sellPriceInCentsForSingleToken = newSellPriceInCents;
	}

    function setBuyPrice(uint256 newBuyPriceInCents) public onlyAdmin {
		buyPriceInCentsForSingleToken = newBuyPriceInCents;
	}

    function setPrivateAddByPublicAdd(address publicAddress, address privateAddress) public onlyAdmin returns (bool) {
        PublicPrivateAddressesMapping[publicAddress] = privateAddress;
        return true;
    }

    function MapEmailIDWithPrivateAdd(string email, address privateAddress) public onlyAdmin returns (bool) {
        var emailInByte = stringToBytes32(email);
        
        EmailIDPrivateAddrMapping[emailInByte] = privateAddress;
        return true;
    }

    function stringToBytes32(string memory source) returns (bytes32 result) {
        assembly {
            result := mload(add(source, 32))
        }
    }

	function mintToken(uint256 mintedAmount) onlyAdmin {
		Balances[admin] += mintedAmount;
		coinSupply += mintedAmount;
        setReservedCoins();
		Transfer(0, this, mintedAmount);
		Transfer(this, admin, mintedAmount);
	}

    function mintICOSupply(uint256 increasedCoins) onlyAdmin {
        uint256 tempICOSupply = ICORemainingBalance + increasedCoins;
        if(Balances[admin] < tempICOSupply) revert();
        ICOSupply += increasedCoins;
        setReservedCoins();
        ICORemainingBalance += increasedCoins;
	}

    function getPrivateAddByPublicAdd(address publicAddr) constant returns (address) {
        return PublicPrivateAddressesMapping[publicAddr];  
    }

    function getPivateAddByEmailID(string email) constant returns (address) {
        var emailInByte = stringToBytes32(email);
        return EmailIDPrivateAddrMapping[emailInByte];
    }

    function getICOSupply() constant returns (uint256) {
        return ICOSupply;
    }

    function getICORemainingBalance() constant returns (uint256) {
        return ICORemainingBalance;
    }

    function getReservedCoins() constant returns (uint256) {
        return reservedCoins;
    }

    function getReservedRemainingCoins() constant returns (uint256) {
        return reservedCoinsBalance;
    }

    function getSellPrice() constant returns (uint256) {
        return sellPriceInCentsForSingleToken;
    }

    function getBuyPrice() constant returns (uint256) {
        return buyPriceInCentsForSingleToken;
    }

    function getFundingTokens() constant returns (uint256) {
        return fundingTokens;
    }

    function getFundingTokensForReserved() constant returns (uint256) {
        return fundingTokensForReserved;
    }

    function getTokenHolders() constant returns (uint256 noOfHolders) {
        return noOfHolders;
    }

    /***********************************************************************************************************/
    //  CROWDSALE LOGIC

    struct Contribution {
        uint amount;                                                                    //amount(in ETH) the person has contributed
        address contributor;
    }

    Contribution[] contributions;

    uint public completedAt;
    uint public priceInUSD;                                                             //price of token (e.g. 1 token = 1 ETH i.e. 10^18 Wei )
    address public creator;
    string campaignUrl;
    byte constant version = 1;

    event LogRefund(address addr, uint amount);
    event LogFundingReceived(address addr, uint amount, uint currentTotal);             //funds received by contributors
    event LogFundingSuccessful(uint totalRaised);                                       //will announce when funding is successfully completed
    event LogFunderInitialized(address creator, string url);

    function SJTCwdSale () {
        creator = admin;
        campaignUrl = "https://stjameshdinvtrust.co/sjt-fund-coin-ico/";                //add SJT website url
        
        LogFunderInitialized(creator, campaignUrl);
    }

    //******************************************************************************/
    //Helps Investors to invest in our CrowdSale

    function releaseTokensByUSD(address _sender, uint256 fundingTokens) onlyAdmin public {
        
        uint256 remainingCoins = getICORemainingBalance();  
        
        if (remainingCoins < fundingTokens) revert();

        contributions.push(
            Contribution({
                amount: fundingTokens,
                contributor: _sender
                })
            );

        releaseTokensForICO(_sender, fundingTokens);
    }
}