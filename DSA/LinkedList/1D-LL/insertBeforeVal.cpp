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
    int x;
    cin >> x;
    Node* wow = new Node(10);
    if(head==nullptr){
        return nullptr;
    }
    if(head->data==x){
        wow->next=head;
        return wow;
    }
    Node* prev= nullptr
    Node* temp=head;
    while(temp){
        cnt++;
        if(cnt==k-1){
             wow->next=prev->next;
             prev->next=wow;
             break;
        }
        prev=temp;
        temp=temp->next;
    }
    return head;
}