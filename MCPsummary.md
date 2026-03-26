# What is MCP in one line
MCP (Model Context Protocol) is an **open-source standard for connecting AI applications to external systems**

# What is standardized by MCP
What MCP standardized can be split into two layer Data Layer, Transport Layer

## Data Layer
The data layer implements a JSON-RPC 2.0 based exchange protocol that defines the message structure and semantics

- Server Feature
- Client Feature
- Utility Feature
- Lifecycle Management

## Transport Layer
The transport layer manages communication channels. There is two options for transport layer 
- Stdio
- Streamable HTTP transport
Transport layer documentation: https://modelcontextprotocol.io/specification/2025-11-25/basic/transports. 
This layer also standardize authc/z between MCP clients and servers
Authorization documentation: https://modelcontextprotocol.io/specification/2025-11-25/basic/authorization


# DataLayer Protocol
This defines three thing
- Primitive (Server and Client feature) 
- Lifecycle
- Notification

## Primitives
Define what clients and servers can offer each other.

### Server Primitives
- Tools: Executable functions that AI applications can invoke to perform actions (e.g., file operations, API calls, database queries)
- Resources: Data sources that provide contextual information to AI applications (e.g., file contents, database records, API responses)
- Prompts: Reusable templates that help structure interactions with language models (e.g., system prompts, few-shot examples)

Client discover these primitives using */list 


### Client Primitives
- Sampling: Allows servers to request language model completions from the client’s AI application.(Server can use LLM)
- Elicitation: Allows servers to request additional information from users. 
- Roots: Allows clients to expose filesystem “roots” to servers
- Logging: Enables servers to send log messages to clients for debugging and monitoring purposes.

## Lifecycle Management
MCP is a stateful protocol Since server/client should know which capability is supported.
This is why MCP requires lifecycle management.

## Notifications
Real time notification between client and server.eg: real time update for tools


# What should be covered by developer to build MCPServer
Allmost everything on Transport Layer is covered by Framework.  
Also Lifecycle Management and Notification, Primitive Discovery is covered by Framework side.  
**User can focus on what Server primitive they want to build and how to use Client Primitive in the implementation.**  
