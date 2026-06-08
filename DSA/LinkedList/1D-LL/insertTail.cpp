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

int main(){
    Node* wow = new Node(10);
    if(head==nullptr){
        return wow;
    }
    Node* temp = head;
    while(temp->next!=null){
        temp=temp->next;
    }
    temp->next=wow;
    return head;
}