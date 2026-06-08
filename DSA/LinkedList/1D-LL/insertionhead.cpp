#include<bits/stdc++.h>
using namespace std;

class Node{
    public:
    int data;
    Node* next;
    Node(int data){
        data = data;
        next= nullptr
    }
};

Node* insertHead(Node* head, int val){
    Node* temp = new Node(val);
    temp->next=head;
    return temp;
}

int main(){
    Node* head = insertHead(head,val);
    cout >> head-> data;
    
}