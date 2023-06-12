#include<string> 
#include<vector> 

using std::vector; 
using std::string;

class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) { 

        string smallest = strs[0];
        string prefix = "";

        // Find the  smallest string
        for(int i = 0; i < strs.size(); i++){ 

            if(strs[i].length() < smallest.length()){ 
                smallest = strs[i];
            }
        } 

        bool prefix_found;
        while(true){ 

            prefix_found = true;

            // Itreate through remaining strings and check if smallest is the largest substring
            for(int i = 0; i < strs.size(); i++){ 

                if(strs[i].rfind(smallest,0) != 0){ 

                    prefix_found = false; 

                    // If not substring, remove last char from string
                    if(smallest.length() > 0){ 
                         smallest.pop_back();  
                         break; 

                    } else { 
                         return prefix;
                    }
                }
            }  

            if(prefix_found){ 
                prefix = smallest; 
                break; 
            }
        } 
        return prefix;   
    }
};