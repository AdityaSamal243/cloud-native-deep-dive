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
    vector<int> arr ={3,2,9.6};
    Node* head = new Node(arr[0]);
    Node* temp = head;
    for(int i=1;i<=3;i++){
        Node* curr=new Node(arr[1]);
        curr->prev=temp;
        temp->next=curr;
        temp=curr;
    }
    return head;
    
}