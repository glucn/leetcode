import java.util.Queue;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

// This is the custom function interface.
// You should not implement it, or speculate about its implementation
interface CustomFunction {
    // Returns f(x, y) for any given positive integers x and y.
    // Note that f(x, y) is increasing with respect to both x and y.
    // i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
    public int f(int x, int y);
};

// Runtime 3 ms
// Memory 35.9 MB

class Solution {
    public List<List<Integer>> findSolution(CustomFunction customfunction, int z) {

      ArrayList<List<Integer>> result = new ArrayList<>();
      boolean[][] visited = new boolean[z*2][z*2];

      Queue<Pair> q = new LinkedList<>();
      q.add(new Pair(1, 1));
      visited[1][1] = true;

      while (!q.isEmpty()) {
        Pair p = q.remove();

        if (customfunction.f(p.i, p.j) == z) {
          result.add(Arrays.asList(p.i, p.j));
        }
        if (customfunction.f(p.i+1, p.j) <= z && !visited[p.i+1][p.j]) {
          q.add(new Pair(p.i+1, p.j));
          visited[p.i+1][p.j] = true;
        }
        if (customfunction.f(p.i, p.j+1) <= z && !visited[p.i][p.j+1]) {
          q.add(new Pair(p.i, p.j+1));
          visited[p.i][p.j+1] = true;
        }
      }
      return result;

    }

  private class Pair {
    int i;
    int j;

    Pair(int i, int j) {
      this.i = i;
      this.j = j;
    }
  }
}
