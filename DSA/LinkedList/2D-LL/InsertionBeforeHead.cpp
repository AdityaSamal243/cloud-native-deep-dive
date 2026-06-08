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
    Node* temp = new Node(10);
    temp->next=head;
    head->prev=temp;
    return temp;
    
}