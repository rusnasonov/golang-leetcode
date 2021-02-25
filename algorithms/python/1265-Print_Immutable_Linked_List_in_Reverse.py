# """
# This is the ImmutableListNode's API interface.
# You should not implement it, or speculate about its implementation.
# """
class ImmutableListNode:
    def printValue(self) -> None: pass
    def getNext(self) -> 'ImmutableListNode': pass

from collections import deque

class Solution:
    def printLinkedListInReverse(self, head: 'ImmutableListNode') -> None:
        nodes = deque()
        while head:
            nodes.append(head)
            head = head.getNext()
        while len(nodes) >  0:
            node = nodes.pop()
            node.printValue()
