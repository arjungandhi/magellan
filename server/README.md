# Server

The Magellan Server is a application that is *currently* designed to be run as a systemd service. 

The server has an api that does the following. 
    - regularly crawls and indexes all documents in a buckets based on their "update frequency/ on change" 
    - provides a place for those documents to be searched
