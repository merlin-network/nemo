<!--
order: 2
-->

# State

## Module Account
The liquid module defines a module account with name `liquid` that has `Minter` and `Burner` module account permissions. The associated bech32 account address is `fury1gggszchqvw2l65my03mak6q5qfhz9cn2y3u50d`. 

## Genesis state

The liquid module does not require any module specific genesis state.

## Store

The liquid module does not store any module specific data. All `bfury` token receipts are minted directly to the delegators account, and the delegation object is transferred to the liquid module account. 