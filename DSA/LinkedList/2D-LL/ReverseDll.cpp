#include<bits/stdc++.h>
using namespace std;

class Node{
    int data;
    Node* next;
    Node* prev;
    Node(int data1){
        data = data1;
        next = nullptr;
        prev = nullptr;
    }
};
int main{
    Node* curr= head;
    Node* last = nullptr;
    while(curr){
        last=curr->prev;
        curr->prev=curr->next;
        curr->next=last;
        curr=curr->prev;
    }
    return last->prev;
    
    
}