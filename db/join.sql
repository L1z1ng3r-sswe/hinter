
Full outer join = outer join, full is a default. Full outer join receives from both crates even if values isn't present, just leave null.

Right outer join / right join: all the values from right table and for them from left -> if doesn't exist in left -> null.

Left outer join / left join: all the values from left table and for them from right -> if doesn't exist in right -> null.

Right inner join and left inner join: do not exist. Inner join retrieves only matching rows from both tables and doesn't depend on which table is left or right.