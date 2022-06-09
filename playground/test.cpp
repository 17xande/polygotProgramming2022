#include <vector>

void print(std::vector<int>& a) {
  printf("vector: ");
  for (const auto& item : a) {
    printf("%d ", item);
  }

  printf("\n");
}

int main() {
  // It's not copying the reference, it's copying values.
  std::vector<int> a;
  std::vector<int> b = a;

  b.push_back(1);

  printf("a: ");
  print(a);
  printf("b: ");
  print(b);
}
