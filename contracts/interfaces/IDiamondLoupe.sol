// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/******************************************************************************\
* Author: Nick Mudge <nick@perfectabstractions.com> (https://twitter.com/mudgen)
* EIP-2535 Diamonds: https://eips.ethereum.org/EIPS/eip-2535
/******************************************************************************/

// A loupe is a small magnifying glass used to look at diamonds.
// These functions look at diamonds
// 查询接口
interface IDiamondLoupe {
    /// These functions are expected to be called frequently
    /// by tools.

    struct Facet {
        address facetAddress; // 本钻石面（Facet）地址
        bytes4[] functionSelectors; // 本钻石面（Facet）所支持的函数选择子的集合
    }

    /// @notice Gets all facet addresses and their four byte function selectors.
    /// @return facets_ Facet // 返回所有的钻石面的结构
    function facets() external view returns (Facet[] memory facets_);

    /// @notice Gets all the function selectors supported by a specific facet.
    /// @param _facet The facet address.
    /// @return facetFunctionSelectors_ // 返回指定钻石面的所有函数选择子
    function facetFunctionSelectors(address _facet) external view returns (bytes4[] memory facetFunctionSelectors_);

    /// @notice Get all the facet addresses used by a diamond.
    /// @return facetAddresses_ 返回所有钻石面的地址
    function facetAddresses() external view returns (address[] memory facetAddresses_);

    /// @notice Gets the facet that supports the given selector.
    /// @dev If facet is not found return address(0).
    /// @param _functionSelector The function selector.
    /// @return facetAddress_ The facet address. //返回指定函数选择子所属的钻石面的地址
    function facetAddress(bytes4 _functionSelector) external view returns (address facetAddress_);
}
