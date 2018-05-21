# Gossip

Connects different nodes over the network with efficient message broadcast and unicast.

&nbsp;
## `Network Requirements`

* Fully connected, every node is connected to every other node.
* Network topology is known in advance and configurable.
  * Assume all services are found on all nodes.
  * Assume single instance of a service on a node.
* All communication is direct.
* No special support for retransmission except standard TCP stack.
* All messages should be signed by nodes and authenticated upon reception. (Peer2Peer)
* Reconnect to topology peers when disconnected with a configurable polling interval.
* Call `OnMessageReceived` when a gossip message is received.
* See inter node message [encoding](../../interfaces/protocol/gossip/json-over-websocket.md).

&nbsp;
## `OnMessageReceived` (method)
> Called when a gossip message is received from another node, forwarded to the services that are subscribed to the message.
* Call `Target.GossipMessageReceived` with message content.
* The target service is responsible to identify the topic and message type and process the message accordingly.

&nbsp;
## `BroadcastMessage` (rpc)
> Broadcasts a message to all the services that are subscribed to the message topic.
* Sends a message to each of the nodes' gossip. (pruning will be added at a later stage)

&nbsp;
## `UnicastMessage` (rpc)
> Sends a message to a single node and service.
* Send message to the specific recipient node.

&nbsp;
## `TopicSubscribe` (rpc)
> Used by services to subscribe to a specific topic. 
Returns a token that can be used in order to unsubscribe.

&nbsp;
## `TopicUnsubscribe` (rpc)
> Used by services to unsubscribe from a specific topic. 
