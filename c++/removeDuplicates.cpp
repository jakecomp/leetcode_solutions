#include<vector>

using std::vector; 
 
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int new_size = 0;
        for(int index = 1; index < nums.size(); index++){
            if(nums[new_size] != nums[index]){ 
                new_size++; 
                nums[new_size] = nums[index];
            }
        }
        return new_size + 1; 
    }
};