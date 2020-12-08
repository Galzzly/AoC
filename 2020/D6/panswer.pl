#!/usr/bin/env perl
$/ = '';

my $part1 = 0;
my $part2 = 0;

while (<>) {
    my $size = scalar split /\n/;

    foreach my $q ('a' .. 'z') {
        my $count =()= m#$q#g;

        $part1++ if ($count);
        $part2++ if ($count == $size);
    }
}

print "Part 1: $part1\n";
print "Part 2: $part2\n";
