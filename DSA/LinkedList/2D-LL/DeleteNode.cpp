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
    int key;
    cin >> key;
    if(head==nullptr) return nullptr;

    Node* temp=head;
    while(temp){
        if(temp->data==key){
            break;
        }
        temp=temp->next;
    }
    Node* back=temp->prev;
    Node* front=temp->next;
    if(front==nullptr){
        back->next=front;
        temp->prev=nullptr;
        free(temp);
        retun head;
    }
    back->next=front;
    front->prev=back;
    temp->next=nullptr;
    temp->back=nullptr;
    delete temp;
    return head;
    
}