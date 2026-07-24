#include<bits/stdc++.h>
using namespace std;

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int t;
    cin >> t;
    while(t--){
        int n;
        cin >> n;
        if(n%3==1 || n%3==2){
            cout << "First\n";
        }
        else{
            cout << "Second\n";
        }
    }
    return 0;
}