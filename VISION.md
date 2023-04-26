# VISION

## Introduction

### Purpose

The purpose of this requirements document is to outline the necessary features, functionality, and specifications for the Nostr Implementation Possibilities (NIPs) monorepo. This document will serve as a reference for contributors, developers, and users who are interested in understanding the scope of the project, its goals, and the expectations for implementing and interacting with the monorepo.

By providing clear and comprehensive requirements, we aim to facilitate effective collaboration, efficient development, and a consistent understanding of the project among all stakeholders. This document will guide the design, implementation, and maintenance of the monorepo to ensure it meets the needs of the Nostr community and adheres to the principles outlined in the NIPs.

### Scope

This requirements document focuses on the development of the go-nostr software, a Nostr-compatible implementation that includes both relay and client functionalities. The software will provide the necessary features and tools for developers and users to interact with the Nostr protocol, allowing for the creation and management of events, as well as communication between clients and relays. The go-nostr software aims to streamline the development of Nostr-compatible applications, promote the adoption of the Nostr network, and support the growth of the decentralized social platform. While not covering the implementation of every single NIP, it will focus on core functionalities and selected NIPs that are essential for a functional and robust Nostr system.

### Definitions, Acronyms, and Abbreviations

TBD

### References

- [RFC 8615: Well-Known Uniform Resource Identifiers (URIs)](https://www.rfc-editor.org/rfc/rfc8615)

### Overvierw

TBD

## Overall Description

### Product Perspective

The go-nostr software is designed as a part of the Nostr network, which is a decentralized and extensible social protocol. The software acts as a bridge between Nostr-compatible applications, providing both relay and client functionalities to enable seamless communication across the network.

### Product Functions

go-nostr serves as a component of the Nostr network by facilitating event creation, management, and communication between clients and relays. It provides core functionalities and supports selected NIPs to streamline the development of Nostr-compatible applications and contribute to the network's growth.

### User Characteristics

The primary users of go-nostr include contributors, developers, and end-users. Contributors are those who actively participate in the development process, while developers leverage the software to create Nostr-compatible applications. End-users, on the other hand, interact with the applications built upon go-nostr to access the Nostr network.

### Constraints

go-nostr must adhere to the specifications outlined in the Nostr protocol and maintain compatibility with other Nostr implementations. Furthermore, it must be efficient, secure, and robust while providing a scalable and user-friendly solution.

| Status | NIP                                                                        |
|--------|----------------------------------------------------------------------------|
|        | NIP-01: Basic protocol flow description                                    |
|        | NIP-02: Contact List and Petnames                                          |
|        | NIP-03: OpenTimestamps Attestations for Events                             |
|        | NIP-04: Encrypted Direct Message                                           |
|        | NIP-05: Mapping Nostr keys to DNS-based internet identifiers               |
|        | NIP-06: Basic key derivation from mnemonic seed phrase                     |
|        | NIP-07: window.nostr capability for web browsers                           |
|        | NIP-08: Handling Mentions --- unrecommended: deprecated in favor of NIP-27 |
|        | NIP-09: Event Deletion                                                     |
|        | NIP-10: Conventions for clients' use of e and p tags in text events        |
|        | NIP-11: Relay Information Document                                         |
|        | NIP-12: Generic Tag Queries                                                |
|        | NIP-13: Proof of Work                                                      |
|        | NIP-14: Subject tag in text events.                                        |
|        | NIP-16: Event Treatment                                                    |
|        | NIP-18: Reposts                                                            |
|        | NIP-19: bech32-encoded entities                                            |
|        | NIP-20: Command Results                                                    |
|        | NIP-21: nostr: URL scheme                                                  |
|        | NIP-22: Event created_at Limits                                            |
|        | NIP-23: Long-form Content                                                  |
|        | NIP-25: Reactions                                                          |
|        | NIP-26: Delegated Event Signing                                            |
|        | NIP-27: Text Note References                                               |
|        | NIP-28: Public Chat                                                        |
|        | NIP-33: Parameterized Replaceable Events                                   |
|        | NIP-36: Sensitive Content                                                  |
|        | NIP-39: External Identities in Profiles                                    |
|        | NIP-40: Expiration Timestamp                                               |
|        | NIP-42: Authentication of clients to relays                                |
|        | NIP-45: Counting results                                                   |
|        | NIP-46: Nostr Connect                                                      |
|        | NIP-50: Keywords filter                                                    |
|        | NIP-51: Lists                                                              |
|        | NIP-56: Reporting                                                          |
|        | NIP-57: Lightning Zaps                                                     |
|        | NIP-58: Badges                                                             |
|        | NIP-65: Relay List Metadata                                                |
|        | NIP-78: Application-specific data                                          |

### Assumptions and Dependencies

The development of go-nostr assumes that the Nostr protocol will continue to evolve and incorporate new features and improvements. As such, the software's design must be flexible enough to accommodate these changes. Additionally, go-nostr depends on the stability and reliability of the Nostr network and its underlying infrastructure.

### Apportioning of Requirements

This document focuses on the essential requirements for go-nostr. However, as the Nostr protocol evolves, new requirements may emerge. To ensure the software remains up-to-date and relevant, these requirements will be allocated to subsequent releases or iterations of the software.

## Specific Requirements

### External Interfaces

#### Software Interfaces

##### API

###### GET /admin/health

Returns the current health status of the service.

###### GET /.well-known/nostr.json?name=local-part

Returns the internet identifier for the given name query parameter.

#### User Interfaces

##### Account Settings Page Component

Allows users to manage their account information and preferences.

##### Add Relay Form Component

Provides a form for users to add a new relay to their account.

##### Appearance Settings Page Component

Enables users to customize the look and feel of the application.

##### Backup Settings Page Component

Facilitates creating and managing backups of user data.

##### Create Short Text Note Event Component

Allows users to create a brief text note within the application.

##### Home Page Component

Displays the main content and navigation options for users.

##### Landing Page Component

Introduces the application and prompts users to sign up or sign in.

##### List Events Component

Displays a list of all events associated with the user's account.

##### List Notifications Component

Shows a list of recent notifications for the user.

##### List Relays Component

Presents a list of all relays the user has added to their account.

##### Network Settings Page Component

Enables users to configure network-related settings and preferences.

##### Not Found Page Component

Informs users that the requested page or resource could not be found.

##### Notifications Page Component

Displays an overview of the user's notifications and settings.

##### Profile Page Component

Shows the user's profile information and allows them to edit it.

##### Remove Relay Form Component

Provides a form for users to remove a relay from their account.

##### Sign In Page Component

Prompts users to enter their credentials to access their account.

##### Sign Out Form Component

Provides a simple form for users to securely log out of their account and end their session.

##### Sign Up Page Component

Allows new users to create an account for the application.

### Functions

#### Add Relay

Allows users to add a new relay to their list of relays, enhancing network connectivity and redundancy.

#### Admin Get Health

Returns the current health status of the service implementing the healthcheck endpoint monitoring pattern.

#### Copy Private Key

Enables users to securely copy their private key to their clipboard for backup or other purposes.

#### Copy Public Key

Enables users to copy their public key to their clipboard for sharing with others or for use in external applications.

#### Get Connection

Get Connection is a product function that establishes a secure client-relay WebSocket connection. This allows clients to communicate with the server in real-time and enables bi-directional data transfer. The connection is encrypted using industry-standard TLS protocols to ensure the security of the communication. Additionally, the connection can be configured to automatically reconnect in case of network disruptions or server downtime, ensuring a reliable connection. This feature is particularly useful for applications that require real-time updates or collaborative features, such as multiplayer games or collaborative editing tools.

#### Get Internet Identifier

Returns the internet identifier for a given name and optionally relays, if enabled. Supports NIP-05.

#### List Events

Displays a list of events, including their details, for the user to browse and interact with.

#### List Notifications

Shows a list of notifications received by the user, allowing them to stay informed about important updates or messages.

#### List Relays

Displays a list of relays currently being used by the user, providing an overview of network connectivity.

#### Publish Message

Enables users to create and publish messages to the network, allowing them to share information and communicate with others.

#### Remove Relay

Allows users to remove a relay from their list of relays, providing control over network connectivity options.

#### Sign In

Enables users to securely log in to their account by providing their credentials.

#### Sign Out

Allows users to securely log out of their account and end their session.

#### Sign Up

Enables new users to create an account by providing the necessary information and agreeing to terms of service.

### Performance Requirements

TBD

### Design Constraints

TBD

### Software System Attributes

TBD

### Other Requirements

TBD
