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
        if(k==1){
            return wow;
        }
        return nullptr;
    }
    if(k==1){
        wow->next=head;
        return wow;
    }
    int cnt=0;
    Node* temp=head;
    while(temp){
        cnt++;
        if(cnt==k-1){
             wow->next=temp->next;
             temp->next=wow;
             break;
        }
        temp=temp->next;
    }
    return head;
}