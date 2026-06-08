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
int lengthDLL(Node* head){
    Node* temp = head;
    int count=0;
    while(temp){
        count++;
        temp=temp->next;
    }
    return count;
}
int main{

    int length=lengthDLL(head);
    if(head==nullptr){
        return nullptr;
    }
    if(length==1){
        return nullptr;
    }
    else{
        Node* temp=head;
        head=head->next;
        head->prev=nullptr;
        temp->next=nullptr;
        free(temp);
        return head;
    }
    
}