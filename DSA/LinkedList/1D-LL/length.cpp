#include<bits/stdc++.h>
using namespace std;

class Node{
    public:
    int data;
    Node* next;

    public:
    Node(int data){
        data = data;
        next = nullptr;
    }
};

Node* ConvertArrLL(vector<int> &arr){
    Node* head = new Node(arr[0]);
    Node* mover = head;
    for(int i=1;i<arr.size();i++){
        Node* temp = new Node(arr[i]);
        mover->next = temp;
        mover = temp;
    }
    return head;
}

int lengthLL(Node* head){
    int count =0;
    Node* temp=head;
    while(temp){
        count++;
        temp=temp->next       
    }
    return count;
}
bool Present(Node* head, int key){
    Node* temp=head;
    while(temp){
        if(temp->data == key){
            return true;
        }
        temp=temp->next;
    }
    return false;
}

int main(){
    vector<int> arr = {1,2,3,4,5};
    Node* head = ConvertArrLL(arr);
    cout << lengthLL(head) << endl;
    cout << Present(head, 3) << endl;
}