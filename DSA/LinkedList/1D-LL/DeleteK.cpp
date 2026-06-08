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


Node* DeleteK(Node* head, int k){
    if(head==nullptr) return nullptr;
    if(k==1){
        Node* temp=head;
        head=head->next
        free(temp);
        return head;
    }
    Node* temp=head;
    Node* prev=nullptr
    int cnt=0;
    while(temp){
        cnt++;
        if(cnt==k){
           // prev->next=prev->next->next;
           prev->next=temp->next
           free(temp);
           break;
        }
        prev=temp;
        temp=temp->next
    }
    return head;
}
int main(){
    int k;
    cin << k;
    Node* head = DeleteK(head,k);
    cout >> head->data

}