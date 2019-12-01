/**
 * // This is Sea's API interface.
 * // You should not implement it, or speculate about its implementation
 * class Sea {
 *     public boolean hasShips(int[] topRight, int[] bottomLeft);
 * }
 */

 // Runtime 1 ms, faster than ?%
 // Memory 30.3 MB, less than ?%

class Solution {
    private Sea sea;

    public int countShips(Sea sea, int[] topRight, int[] bottomLeft) {
        this.sea = sea;

        return dfs(topRight, bottomLeft);
    }
    
    private int dfs(int[] topRight, int[] bottomLeft) {
        if (!this.sea.hasShips(topRight, bottomLeft)) {
            return 0;
        }
        
        int width = topRight[0] - bottomLeft[0];
        int height = topRight[1] - bottomLeft[1];
        if (width == 0 && height == 0) {
            return 1;
        }
        
        if (width > height) {
            int mid = (bottomLeft[0] + topRight[0]) / 2;
            int leftCount = dfs(new int[]{mid, topRight[1]}, bottomLeft);
            int rightCount = dfs(topRight, new int[]{mid+1, bottomLeft[1]});
            return leftCount + rightCount;
        }
        
        int mid = (bottomLeft[1] + topRight[1]) / 2;
        int bottomCount = dfs(new int[]{topRight[0], mid}, bottomLeft);
        int topCount = dfs(topRight, new int[]{bottomLeft[0], mid+1});
        return bottomCount + topCount;
    }
}