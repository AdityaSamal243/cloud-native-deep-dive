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
    if(k==1){
        Node* temp= head;
        head = head->next;
        head->prev=nullptr;
        temp->next=nullptr;
        delete temp;
    }
    Node* temp=head;
    int count==0;
    while(temp){
        count++;
        if(count==k){
             break;
        }
        temp=temp->next;
    }
    back= temp->prev;
    ahead=temp->next;
    back->next=ahead;
    if(ahead == nullptr){
        back->next=ahead;
        temp->prev=nullptr;
        free(temp);
        return head;
    }
    back->next=ahead;
    ahead->prev=back;
    temp->next=nullptr;
    temp->prev=nullptr;
    return head;
    
    
}