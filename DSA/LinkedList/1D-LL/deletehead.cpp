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

Node* deletehead(Node* head){
    if(head ==nullptr) return head;
    Node* temp = head;
    head  = head->next;
    free(temp);
    return head;
}

int main(){

    Node* head = deletehead(head)
    
}