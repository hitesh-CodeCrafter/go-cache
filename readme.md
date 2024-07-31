# Overview

This project implements a basic cache using a combination of a doubly linked list and a hash map. The cache is designed to efficiently manage the insertion, retrieval, and eviction of elements based on their usage. The least recently used (LRU) element is removed when the cache reaches its maximum size.
Features

    Efficient Cache Operations: Quick access, insertion, and deletion of elements.
    LRU Eviction Policy: Automatically evicts the least recently used element when the cache exceeds its size limit.
    Doubly Linked List: Utilizes a doubly linked list to manage the order of elements.
    Hash Map: Uses a hash map to enable O(1) access to cache elements.

## Implementation
Data Structures

    Node: Represents an element in the cache. Contains a value and pointers to the previous and next nodes.
    Queue: A doubly linked list that maintains the order of elements in the cache. The head and tail nodes are dummy nodes to simplify edge cases.
    Cache: Contains the hash map and the queue. Manages the insertion, retrieval, and eviction of elements.

## Functions

    NewCacheCreation: Initializes a new cache with an empty queue and hash map.
    NewQueue: Initializes a new queue with dummy head and tail nodes.
    Check: Checks if an element is in the cache. If it is, it moves the element to the front of the queue. If it is not, it adds the element to the cache.
    Remove: Removes an element from the queue and hash map.
    Add: Adds an element to the front of the queue and updates the hash map. Evicts the least recently used element if the cache size exceeds the limit.
    Display: Displays the current state of the cache.
