/*******************************************************************************
  Title          : validate.c
  Author         : kkhan01
  Description    : AOC 2019 Day 4
  Purpose        :
  Usage          : mpirun -np <num_cores> validate
  Build with     : mpicc -o validate validate.c
  Acknowledgments:
  License        :
*******************************************************************************/
#include "mpi.h"

#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>

#define ROOT 0

void validate(long long int id, int p, int lower_bound, int upper_bound,
              int *part1, int *part2);
/** validate
 * @id: process unique id
 * @p: number of processors
 * @lower_bound: starting number for loop
 * @upper_bound: ending number for loop
 * @part1: store answer for part 1
 * @part2: store answer for part 2
 * AOC 2019 Day 4
 * void due to using reference values
 */

int main(int argc, char *argv[]) {
  int id; /* rank of executing process */
  int p;  /* number of processes */

  // problem specific nums
  int lower_bound = 382345; // 382345
  int upper_bound = 843167; // 843167

  // accumulators
  int local_part1 = 0;
  int local_part2 = 0;
  int global_part1 = 0;
  int global_part2 = 0;

  // too lazy to error check args

  MPI_Init(&argc, &argv);
  MPI_Comm_rank(MPI_COMM_WORLD, &id);
  MPI_Comm_size(MPI_COMM_WORLD, &p);

  validate(id, p, lower_bound, upper_bound, &local_part1, &local_part2);

  // 2 reductions because why malloc :shrug:
  MPI_Reduce(&local_part1, &global_part1, 1, MPI_INT, MPI_SUM, 0,
             MPI_COMM_WORLD);
  MPI_Reduce(&local_part2, &global_part2, 1, MPI_INT, MPI_SUM, 0,
             MPI_COMM_WORLD);

  if (ROOT == id) {
    printf("PART 1: %d\nPART 2: %d\n", global_part1, global_part2);
    fflush(stdout);
  }

  MPI_Finalize();
  return 0;
}

short valid1(int num) {
  // somewhat hard coded, but aoc did say length 6
  char s_num[7];
  sprintf(s_num, "%d", num);
  s_num[6] = '\0';

  // check if two pairs are equal
  int repeat = 0;

  // check if sorted
  for (int i = 1; i < 6; ++i) {
    if (s_num[i - 1] > s_num[i]) {
      return 0;
    }

    if (s_num[i - 1] == s_num[i]) {
      repeat = 1;
    }
  }

  return 1 && repeat;
}

short valid2(int num) { return 0; }

void validate(long long int id, int p, int lower_bound, int upper_bound,
              int *part1, int *part2) {
  for (int i = lower_bound + id; i < upper_bound; i += p) {
    // part 1
    *part1 += valid1(i);

    // part 2
    *part2 += valid2(i);
  }
}
