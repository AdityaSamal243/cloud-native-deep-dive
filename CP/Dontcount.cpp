#include<bits/stdc++.h>
using namespace std;

int main(){
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    int t;
    cin >> t;
    while(t--){
        int n,m;
        cin >> n >> m;
        string x;
        cin >> x;
        string s;
        cin >> y;
        int count = 0;
        int flag =0;
        if(x == s){
            cout << 0 << '\n';
        }
        else if(x.length()>=s.length()){
              x = x + x;
              if(x.contains(s)){
                cout << 1 << '\n';
              }
              else{
                cout << -1 << '\n';
              }
        }
        else{
             while(x.length()<=s.length()){
                if(x.contains(s)){
                    flag =1;
                    cout << count << '\n';
                    break;
                }
                x = x + x;
                count++;
            }
            if(flag==0){
                if(x.contains(s)){
                  cout << count << '\n';
                }
                else{
                  cout << -1 << '\n';
                }
        }
        }
    }
    return 0;
}