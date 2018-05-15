#include <cstdio>
#include <cstdlib>
#include <algorithm>

int main(int argc, char *args[]) {
  if (argc != 3) {
    printf("Usage: %s N n\n"
        "N - total number of balls\n"
        "n - number of trials\n", args[0]);
    exit(1);
  }
  int N, n;

  N = atoi(args[1]);
  n = atoi(args[2]);

  int *B; // balls
  B = (int*) malloc(N*sizeof(int));

  for (int i = 0; i < N; ++i) B[i] = i;

  int *T; // trials
  T = (int*) calloc(n, sizeof(int));

  int *X; // random variable
  X = (int*) calloc(N+1, sizeof(int));

  for (int i = 0; i < n; ++i) {
    std::random_shuffle(B, B+N);
    printf("Sequence order:");
    for (int j = 0; j < N; ++j) printf(" %d", B[j]);
    putchar('\n');

    for (int j = 0; j < N; ++j)
      if (j == B[j])
        ++T[i];
    ++X[T[i]];
  }
  int t = 0;
  for (int i = 0; i < N+1; ++i) t += X[i];

  printf("\nTotal pairings per trial:\n");
  for (int i = 0; i < n; ++i) printf(" n=%2d", i);
  putchar('\n');
  double m = 0;
  for (int i = 0; i < n; ++i) {
    printf(" %4d", T[i]);
    m += T[i];
  }
  putchar('\n');

  printf("\nRandom variable distribution:\n");
  for (int i = 0; i < N+1; ++i) printf(" X=%3d", i);
  putchar('\n');
  for (int i = 0; i < N+1; ++i) printf(" %5d", X[i]);
  putchar('\n');
  for (int i = 0; i < N+1; ++i) printf(" %.3f", double(X[i])/t);
  putchar('\n');
  printf("\nEstimated expected value of X: %f\n", m/double(n));

  return 0;
}
