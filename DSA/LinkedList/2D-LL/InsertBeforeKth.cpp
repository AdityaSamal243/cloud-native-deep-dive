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
    int k;
    cin >> k;
    Node* newNode = new Node(10);
    if(head==nullptr){
            return nullptr;
    }
    if(k==1){
            newNode->next=head;
            head->prev=newNode;
            return newNode;
    }
    Node* temp=head;
    int count=0;
    while(temp){
        count++;
        if(count==k){
            Node* back = temp->prev;
            back->next=newNode;
            newNode->next=temp;
            temp->prev=newNode;
            newNode->prev=back;
            return head;
           
        }
        temp=temp->next;
    }
}