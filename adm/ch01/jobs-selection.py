#!/usr/bin/env python3


# Algorithm to select the max number of non-overlapping jobs,
# given their start and end times.
def select_max_jobs(jobs):
    m = []  # The max number of jobs selected.  To be returned.
    jobs = sorted(jobs)

    i = 0
    while i < len(jobs) - len(m):
        s = []  # Jobs selected in this iteration.
        j = i
        while j < len(jobs):
            if len(s) == 0 or s[-1][1] <= jobs[j][0]:
                s.append(jobs[j])
            j += 1
        if len(m) < len(s):
            m = s
        i += 1
    return m


if __name__ == '__main__':
    jobs = [(1,9), (4,5), (5,7), (6,9), (2,8), (3,4),
            (9,13), (15,17), (17,18), (18,20)]
    print(select_max_jobs(jobs))
