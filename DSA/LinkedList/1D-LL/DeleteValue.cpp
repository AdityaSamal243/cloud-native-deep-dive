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

Node* DeleteValue(Node* head, int val){
    if(head==nullptr) return nullptr;
    //check if value is first element 
    //if yeas delete head
    if(head->data==val){
        Node* temp=head;
        head=head->next;
        free(temp);
        return head;
    }
    Node* temp= head->next;
    Node* prev = nullptr;
    while(temp){
        if(temp->data==val){
            prev->next=prev->next->next;
            free(temp);
            break;
        }
        prev=temp;
        temp=temp->next;
    }
    return head;
}

int main(){
    Node* head = DeleteValue(head,val);
}