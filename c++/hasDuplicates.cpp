#include<vector>
#include<map> 

using std::vector; 

class Solution {
public:
    bool containsDuplicate(vector<int>& nums) { 

      std::map<int,int> duplicates;

      for(int index = 0;index < nums.size(); index++){ 

        if(duplicates.find(nums[index]) != duplicates.end()){ 

          return true;
        } 
        duplicates[nums[index]] = 1;
      }
      return false; 
    }
};