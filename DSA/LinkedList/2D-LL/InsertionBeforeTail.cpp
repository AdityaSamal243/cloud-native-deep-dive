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
    Node* newNode = new Node(10);
    
    if(head->next==null){
        newNode->next=head;
        head->prev=newnode;
    }

    Node* temp = head;
    while(temp->next!=null){
        temp=temp->next;
    }
    Node* back = temp->prev;
    newNode->next= temp;
    temp->prev= newNode;
    newNode->prev=back;
    back->next= newNode;
    return head;
    
} 