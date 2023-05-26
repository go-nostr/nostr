# REQUIREMENTS

## Introduction

### Purpose

This document defines the necessary features, functionality, and specifications for the go-nostr/nostr monorepo. It is intended for contributors, developers, and users who wish to understand the scope and objectives of the project, and the expectations for interacting with the monorepo.

This document aims to:

1. Facilitate effective collaboration among stakeholders
2. Enhance development efficiency
3. Foster a shared understanding of the project
4. Guide the design, implementation, and maintenance of the monorepo

The requirements herein are intended to ensure the monorepo satisfies the Nostr community's needs and adheres to the principles outlined in the Nostr Improvement Proposals (NIPs).

### Scope

This document outlines the requirements for the go-nostr/nostr software, a Nostr-compatible implementation that encompasses both relay and client functionalities.

The go-nostr/nostr software:

- The system MUST provide essential features and tools for developers and users to interact with the Nostr protocol
- The system MUST enable the creation and management of events
- The system MUST facilitate communication between clients and relays
- The system SHOULD streamline the development of Nostr-compatible applications
- The system SHOULD promote the adoption of the Nostr network
- The system SHOULD support the growth of a decentralized social platform

However, the go-nostr/nostr software:

- The system MAY NOT implement every single Nostr Improvement Proposal (NIP)
- The system SHOULD Focus on core functionalities and selected NIPs crucial for a functional and robust Nostr system

The aim of the go-nostr/nostr software is to enhance the user experience and expand the capabilities of the Nostr network. This aligns with the broader goals of the Nostr community and the guidelines outlined in higher-level specifications.

### Definitions, Acronyms, and Abbreviations

- **Monorepo**: A monolithic repository that stores all of an organization's deliverable code in a single repository.
- **Nostr**: A decentralized social protocol.
- **NIP**: Nostr Improvement Proposal, a design document providing information to the Nostr community, describing new features or processes.
- **Relay**: A server that passes data from a source to a destination in a network.
- **Client**: A computer program that sends requests to servers in a network for services to be performed.

For further definitions and abbreviations related to the Nostr protocol, please refer to the [Nostr Glossary](https://github.com/nostr-protocol/nips).

### References

- [Nostr Improvement Possibilities (NIPs)](https://github.com/nostr-protocol/nips)
- [RFC 8615: Well-Known Uniform Resource Identifiers (URIs)](https://www.rfc-editor.org/rfc/rfc8615)

### Overview

The go-nostr/nostr software is a Nostr-compatible implementation that provides both relay and client functionalities. It is designed to interact with the Nostr protocol and facilitates the creation, management of events, and communication between clients and relays. The purpose of this software is to simplify the development process of Nostr-compatible applications, promote the adoption of the Nostr network, and support the growth of the decentralized social platform.

This requirements document provides a comprehensive list of features, functionalities, and specifications that guide the design, development, and maintenance of the go-nostr/nostr software. It serves as a reference for stakeholders including contributors, developers, and users, fostering a shared understanding of the project's objectives and expectations.

While not covering every single Nostr Improvement Proposal (NIP), the focus of the go-nostr/nostr software is on the implementation of core functionalities and selected NIPs essential for a functional and robust Nostr system. This approach aligns with the broader goals of the Nostr community and adheres to the principles outlined in higher-level specifications.

The requirements are set forth in Behavior-Driven Development (BDD) style to:

- **MUST**: define the mandatory requirements.
- **SHOULD**: specify the recommended but not obligatory requirements.
- **MAY**: Provide optional requirements.

## Overall Description

### Product Perspective

The go-nostr/nostr software is a component of the Nostr network, a decentralized and extensible social protocol. It acts as a conduit for Nostr-compatible applications, offering both relay and client functionalities for seamless network communication.

### Product Functions

As a key component of the Nostr network, go-nostr/nostr:

- The system MUST facilitate event creation and management.
- The system MUST enable communication between clients and relays.
- The system SHOULD provide core functionalities.
- The system MAY support selected Nostr Improvement Proposals (NIPs) to enhance the development of Nostr-compatible applications and stimulate network growth.

### User Characteristics

go-nostr/nostr serves:

- **Contributors**: Individuals actively involved in the development process.
- **Developers**: Users who utilize the software to build Nostr-compatible applications.
- **End-users**: Individuals who interact with applications built using go-nostr/nostr to access the Nostr network.

### Constraints

go-nostr/nostr adheres to the following constraints:

- The system MUST adhere to the Nostr protocol specifications.
- The system MUST maintain compatibility with other Nostr implementations.
- The system SHOULD be efficient, secure, and robust.
- The system SHOULD provide a scalable and user-friendly solution.

### Assumptions and Dependencies

The development of go-nostr/nostr assumes:

- The Nostr protocol will continue to evolve, incorporating new features and improvements.
- The software design will be flexible to accommodate changes.
- The software depends on the stability and reliability of the Nostr network and its underlying infrastructure.

### Apportioning of Requirements

This document emphasizes essential requirements for go-nostr/nostr. As the Nostr protocol evolves, new requirements may arise. These will be allocated to future releases or iterations of the software to ensure its continuous relevance and update.

## Specific Requirements

### External Interfaces

#### System Interfaces

##### Get Well Known Internet Identifier Endpoint

The system MUST return the internet identifier corresponding to the provided name query parameter.

##### Establish Websocket Connection Endpoint

The system MUST establish a Websocket connection between the client and relay to enable real-time data transmission.

#### Software Interfaces

##### Basic Event Kinds

The system MUST support the `set_metadata`, `text_note`, and `recommend_server` basic event kinds.

##### List Filter Attributes

The system SHOULD support filter attributes containing lists (such as `ids`, `kinds`, or `#e`) as JSON arrays with one or more values.

##### IDs and Authors List

The system SHOULD support `ids` and `authors` lists containing lowercase hexadecimal strings. These may either be an exact 64-character match or a prefix of the event value.

##### Filter Limit

The system SHOULD Use the `limit` property of a filter for the initial query only. This can be ignored subsequently.

##### Multiple Filters in Request Message

The system MAY support that a `REQ` message can contain multiple filters. In this case, events that match any of the filters are to be returned (multiple filters are interpreted as `||` conditions).

##### Event Schema

The system MUST support the following event schema:

```plain
{
  "id": <32-bytes lowercase hex-encoded sha256 of the serialized event data>,
  "pubkey": <32-bytes lowercase hex-encoded public key of the event creator>,
  "created_at": <unix timestamp in seconds>,
  "kind": <integer>,
  "tags": [
    ["e", <32-bytes hex of the id of another event>, <recommended relay URL>],
    ["p", <32-bytes hex of a pubkey>, <recommended relay URL>],
    ... // other kinds of tags may be included later
  ],
  "content": <arbitrary string>,
  "sig": <64-bytes hex of the signature of the sha256 hash of the serialized event data, which is the same as the "id" field>
}
```

##### Event ID Calculation

The system MUST support obtaining the event id by serializing the event and hashing it using sha256.

##### Filter Matching

The system MUST support obtaining the event id by serializing the event and hashing it using sha256.

##### Subscription ID

The system SHOULD support subscription_id as an arbitrary, non-empty string of max length 64 characters, which should be used to represent a subscription.

##### Filter Schema

The system SHOULD support filters as a JSON object that determines what events will be sent in that subscription:

```plain
{
  "ids": <a list of event ids or prefixes>,
  "authors": <a list of pubkeys or prefixes, the pubkey of an event must be one of these>,
  "kinds": <a list of a kind numbers>,
  "#e": <a list of event ids that are referenced in an "e" tag>,
  "#p": <a list of pubkeys that are referenced in a "p" tag>,
  "since": <an integer unix timestamp, events must be newer than this to pass>,
  "until": <an integer unix timestamp, events must be older than this to pass>,
  "limit": <maximum number of events to be returned in the initial query>
}
```

#### User Interfaces

##### Account Settings Page Component

The system MUST enable users to manage their account information and preferences.

##### Add Relay Form Component

The system MUST provide a form for users to add a new relay to their account.

##### Appearance Settings Page Component

The system MUST enable users to customize the application's look and feel.

##### Backup Settings Page Component

The system MUST facilitate creating and managing backups of user data.

##### Create Short Text Note Event Component

The system MUST allow users to create a brief text note within the application.

##### Home Page Component

The system MUST display the main content and navigation options for users.

##### Landing Page Component

The system MUST introduce the application and prompt users to sign up or sign in.

##### List Events Component

The system MUST display a list of all events associated with the user's account.

##### List Notifications Component

The system MUST show a list of recent notifications for the user.

##### List Relays Component

The system MUST present a list of all relays the user has added to their account.

##### Network Settings Page Component

The system MUST enable users to configure network-related settings and preferences.

##### Not Found Page Component

The system MUST inform users that the requested page or resource could not be found.

##### Notifications Page Component

The system MUST display an overview of the user's notifications and settings.

##### Profile Page Component

The system MUST show the user's profile information and allow them to edit it.

##### Remove Relay Form Component

The system MUST provide a form for users to remove a relay from their account.

##### Send Event Message Command

The system MUST provide a command for users to send event messages within the application.

##### Send Request Message Command

The system MUST provide a command for users to send request messages within the application.

##### Send Close Message Command

The system MUST provide a command for users to send close messages within the application.

##### Sign In Page Component

The system MUST prompt users to enter their credentials to access their account.

##### Sign Out Form Component

The system MUST provide a simple form for users to securely log out of their account and end their session.

##### Sign Up Page Component

The system MUST allow new users to create an account for the application.

### Functions

#### Close Connection on Close Message

The system SHOULD close the connection upon receiving the CLOSE message with the same subscription_id.

#### Encode Event Signature

The system SHALL encode signatures, and public key according to the Schnorr signatures standard for the curve secp256k1.

#### Overwrite Subscription on Same Subscription ID

The system SHOULD overwrite the previous subscription upon receiving the REQ message with the same subscription_id.

#### Query Database on Request Message

The system SHOULD query the relay internal database upon receiving a REQ message, return events that match the filter, store that filter, and send again all future events it receives to that same websocket until the websocket is closed.

#### Send Event Message from Client to Relay

The system SHALL be able to send "EVENT" messages from client to relay and publish events.

#### Send Request Message from Client to Relay

The system SHALL be able to send "REQ" messages from client to relay and create subscriptions.

#### Send Close Message from Client to Relay

The system SHALL be able to send "CLOSE" messages from client to relay and stop previous subscriptions.

#### Send Event Message from Relay to Client

The system SHALL be able to send "EVENT" messages from relay to client and send events requested by clients.

#### Send End of Stored Events Message from Relay to Client

The system SHALL be able to send "EOSE" messages from relay to client and indicate the end of stored events.

#### Send Notice Message from Relay to Client

The system SHALL be able to send "NOTICE" messages from relay to client and send human-readable error messages to the client.

### Performance Requirements

#### Single Websocket Connection

The system SHOULD NOT open more than one websocket connection between client and relay.

### Design Constraints

#### NIPs

| Support       | NIP                                                                                                                                    |
|---------------|----------------------------------------------------------------------------------------------------------------------------------------|
| Supported | [NIP-01: Basic protocol flow description](https://github.com/nostr-protocol/nips/blob/master/01.md)                                    |
| Not Supported | [NIP-02: Contact List and Petnames](https://github.com/nostr-protocol/nips/blob/master/02.md)                                          |
| Not Supported | [NIP-03: OpenTimestamps Attestations for Events](https://github.com/nostr-protocol/nips/blob/master/03.md)                             |
| Not Supported | [NIP-04: Encrypted Direct Message](https://github.com/nostr-protocol/nips/blob/master/04.md)                                           |
| Not Supported | [NIP-05: Mapping Nostr keys to DNS-based internet identifiers](https://github.com/nostr-protocol/nips/blob/master/05.md)               |
| Not Supported | [NIP-06: Basic key derivation from mnemonic seed phrase](https://github.com/nostr-protocol/nips/blob/master/06.md)                     |
| Not Supported | [NIP-07: window.nostr capability for web browsers](https://github.com/nostr-protocol/nips/blob/master/07.md)                           |
| Not Supported | [NIP-08: Handling Mentions --- unrecommended: deprecated in favor of NIP-27](https://github.com/nostr-protocol/nips/blob/master/08.md) |
| Not Supported | [NIP-09: Event Deletion](https://github.com/nostr-protocol/nips/blob/master/09.md)                                                     |
| Not Supported | [NIP-10: Conventions for clients' use of e and p tags in text events](https://github.com/nostr-protocol/nips/blob/master/10.md)        |
| Not Supported | [NIP-11: Relay Information Document](https://github.com/nostr-protocol/nips/blob/master/11.md)                                         |
| Not Supported | [NIP-12: Generic Tag Queries](https://github.com/nostr-protocol/nips/blob/master/12.md)                                                |
| Not Supported | [NIP-13: Proof of Work](https://github.com/nostr-protocol/nips/blob/master/13.md)                                                      |
| Not Supported | [NIP-14: Subject tag in text events](https://github.com/nostr-protocol/nips/blob/master/01.14).                                        |
| Not Supported | [NIP-16: Event Treatment](https://github.com/nostr-protocol/nips/blob/master/16.md)                                                    |
| Not Supported | [NIP-18: Reposts](https://github.com/nostr-protocol/nips/blob/master/18.md)                                                            |
| Not Supported | [NIP-19: bech32-encoded entities](https://github.com/nostr-protocol/nips/blob/master/19.md)                                            |
| Not Supported | [NIP-20: Command Results](https://github.com/nostr-protocol/nips/blob/master/20.md)                                                    |
| Not Supported | [NIP-21: nostr: URL scheme](https://github.com/nostr-protocol/nips/blob/master/21.md)                                                  |
| Not Supported | [NIP-22: Event created_at Limits](https://github.com/nostr-protocol/nips/blob/master/22.md)                                            |
| Not Supported | [NIP-23: Long-form Content](https://github.com/nostr-protocol/nips/blob/master/23.md)                                                  |
| Not Supported | [NIP-25: Reactions](https://github.com/nostr-protocol/nips/blob/master/25.md)                                                          |
| Not Supported | [NIP-26: Delegated Event Signing](https://github.com/nostr-protocol/nips/blob/master/26.md)                                            |
| Not Supported | [NIP-27: Text Note References](https://github.com/nostr-protocol/nips/blob/master/27.md)                                               |
| Not Supported | [NIP-28: Public Chat](https://github.com/nostr-protocol/nips/blob/master/28.md)                                                        |
| Not Supported | [NIP-33: Parameterized Replaceable Events](https://github.com/nostr-protocol/nips/blob/master/33.md)                                   |
| Not Supported | [NIP-36: Sensitive Content](https://github.com/nostr-protocol/nips/blob/master/36.md)                                                  |
| Not Supported | [NIP-39: External Identities in Profiles](https://github.com/nostr-protocol/nips/blob/master/39.md)                                    |
| Not Supported | [NIP-40: Expiration Timestamp](https://github.com/nostr-protocol/nips/blob/master/40.md)                                               |
| Not Supported | [NIP-42: Authentication of clients to relays](https://github.com/nostr-protocol/nips/blob/master/42.md)                                |
| Not Supported | [NIP-45: Counting results](https://github.com/nostr-protocol/nips/blob/master/45.md)                                                   |
| Not Supported | [NIP-46: Nostr Connect](https://github.com/nostr-protocol/nips/blob/master/46.md)                                                      |
| Not Supported | [NIP-50: Keywords filter](https://github.com/nostr-protocol/nips/blob/master/50.md)                                                    |
| Not Supported | [NIP-51: Lists](https://github.com/nostr-protocol/nips/blob/master/51.md)                                                              |
| Not Supported | [NIP-56: Reporting](https://github.com/nostr-protocol/nips/blob/master/56.md)                                                          |
| Not Supported | [NIP-57: Lightning Zaps](https://github.com/nostr-protocol/nips/blob/master/57.md)                                                     |
| Not Supported | [NIP-58: Badges](https://github.com/nostr-protocol/nips/blob/master/58.md)                                                             |
| Not Supported | [NIP-65: Relay List Metadata](https://github.com/nostr-protocol/nips/blob/master/65.md)                                                |
| Not Supported | [NIP-78: Application-specific data](https://github.com/nostr-protocol/nips/blob/master/78.md)                                          |****

### Software System Attributes

TBD

### Other Requirements

TBD
