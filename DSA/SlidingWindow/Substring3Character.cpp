//Given a string s consisting only of characters a, b and c.
//Return the number of substrings containing at least one occurrence of all these characters a, b and c.

//s="abcabc" 
//output=10 , abc,abca,abcab.abcabc...............


// better
class Solution {
public:
    int numberOfSubstrings(string s) {
         int l=0;
         int r=0;
         int n= s.length();
         int count=0;
         unordered_map<char,int> um;
         while(r<n){
            um[s[r]]++;
            while(um.size()==3){
                count+=n-r;
                um[s[l]]--;
                if(um[s[l]]==0) um.erase(s[l]);
                l++;
            }
            r++;
         }
         return count;
    }
};


//optimal

class Solution {
public:
    int numberOfSubstrings(string s) {
         int n= s.length();
         int count=0;
         vector<int> lastseen(3,-1);
         for(int i=0;i<n;i++){
            lastseen[s[i]-'A']=i;
            if(lastseen[0]!=-1 && lastseen[1]!=-1 && lastseen[2]!=-1){
                count+=1+min(lastseen[2],min(lastseen[0],lastseen[1]));
            }
         }
         return count;
    }
};