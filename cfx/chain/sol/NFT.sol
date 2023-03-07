// SPDX-License-Identifier: MIT
pragma solidity =0.8.7;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract MetaderbyHorse is ERC721, Ownable {

    address private adminAddr = 0x1396d5e28Dd913970C4dB097e086829181aA39EC;
    address private backendSignAddr = 0x1396d5e28Dd913970C4dB097e086829181aA39EC;

    string public _baseUrl;


    constructor()
    ERC721("Metaderby", "Derby")
    {

        // initial base url
        _baseUrl = "https://d2vvute2v3y7pn.cloudfront.net/NFT/ticket/";

    }

    modifier adminUser {
        //require(msg.sender == adminAddr, "msg.sender is not admin user");
        _;
    }

    modifier backendUser {
        //require(msg.sender == backendSignAddr, "msg.sender is not backend user");
        _;
    }

    function _baseURI()
    internal
    view
    virtual
    override
    returns (string memory)
    {
        return _baseUrl;
    }

    function setBaseUrl(string memory baseUrl)
    external
    backendUser
    {
        _baseUrl = baseUrl;
    }

    function setBackendAddress(address addr)
    external
    adminUser
    returns (bool)
    {
        require(addr != address(0), "address invalid");

        backendSignAddr = addr;

        return true;
    }


    function recover(bytes32 hash, bytes memory sig)
    internal
    pure
    returns (address)
    {
        bytes32 r;
        bytes32 s;
        uint8 v;

        //Check the signature length
        if (sig.length != 65) {
            return (address(0));
        }

        // Divide the signature in r, s and v variables
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }

        // Version of signature should be 27 or 28, but 0 and 1 are also possible versions
        if (v < 27) {
            v += 27;
        }

        // If the version is correct return the signer address
        if (v != 27 && v != 28) {
            return (address(0));
        } else {
            return ecrecover(hash, v, r, s);
        }
    }

    function validUser(uint256 num, uint256 timestamp, uint256 amount, bytes memory sign)
    public
    view
    returns (bool)
    {
        require(timestamp + 120 > block.timestamp, "timestamp <= block.timestamp");
        require(timestamp < block.timestamp, "timestamp > block.timestamp");

        address to = msg.sender;
        bytes memory s = abi.encodePacked(num, timestamp, amount, to);
        bytes32 hash = keccak256(s);

        address addr = recover(hash, sign);
        return addr == backendSignAddr;
    }

    function _doMint(uint256 num, address to)
    internal
    {

        require(!_exists(num), "tokenId is existent");

        _safeMint(to, num);

    }


    function mint(
        uint256 num,
        uint256 timestamp,
        uint256 amount,
        bytes memory sign
    )
    external
    payable
    {
        bool valid = validUser(num, timestamp, amount, sign);
        require(valid, "user invalid");

        require(amount <= msg.value, "value < amount");
        payable(backendSignAddr).transfer(msg.value);

        _doMint(num, msg.sender);
    }

    function mintForUser(uint256 num, address to)
    external
    backendUser
    {
        _doMint(num, to);
    }

    function burn(uint256 tokenId)
    external
    {
        address owner = ownerOf(tokenId);
        require(owner == msg.sender, "owner is invalid");
        _burn(tokenId);
    }

}