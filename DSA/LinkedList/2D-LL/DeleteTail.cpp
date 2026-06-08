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
    

    if(head==nullptr) return nullptr;
    if(head->next==nullptr) return nullptr;

    Node* temp = head;
    while(temp->next!=null){
        
        temp=temp->next;
    }
    back=temp->prev;
    back->next=nullptr;
    temp->prev=nullptr;
    delete temp;
    return head;
    
    
}