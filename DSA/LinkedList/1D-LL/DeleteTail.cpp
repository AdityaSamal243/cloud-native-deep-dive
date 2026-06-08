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

Node* DeleteTail(Node* head){

    if(head==nullptr || head->next==nullptr) return nullptr

    Node* temp=head;
    NOde* prev;
    while(temp->next!=nullptr){
        prev=temp;
        temp=temp->next;
    }
    prev->next= nullptr;
    free(temp);
    return head;

}

int main(){
    Node* head = DeleteTail(head);
    cout << head->data << endl;
}