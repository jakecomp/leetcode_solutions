class Solution {
public:
    double myPow(double x, int n) { 

        if(n == 0 && x != 0){ 
            return double(1);
        } 
        
        double result;
        if(n > 0){ 
            result = x;
            for(int i = 2;i <= n;i++){
                result *= x;
            }

        } else { 
            result = 1/x;
            int positive_n = -1 * n; 
            for(int i = 2;i <= positive_n;i++){ 
                result /= x;
            }
        }
        return result;
        
    } 
};