package com.dalen.aoc

class Day11 {

    fun solve1(): Long {
        return solve(20, 3)
    }

    fun solve2(): Long {
        return solve(10000, 1)
    }

    private fun solve(loops: Int, relief: Int): Long {
        val monkeys = initMonkeys()
        val lcm = monkeys.map { it.divisible }.reduceRight {a,b -> a * b}
        val inspections = mutableMapOf<Int, Long>()

        for (monkey in monkeys) {
            inspections[monkey.name] = 0
        }

        for (round in (0 until loops)) {
            for (monkey in monkeys) {
                for (item in monkey.items) {
                    val p = doItem(monkey, item, inspections, relief, lcm)
                    monkeys[p.first].items.add(p.second)
                }
                monkey.items.removeAll(monkey.items)
            }
        }

        println(inspections)
        val top2 = inspections.values.sorted().reversed().take(2)
        return top2[0] * top2[1]
    }

    private fun doItem(
        monkey: Monkey,
        item: Long,
        inspections: MutableMap<Int, Long>,
        relief: Int,
        lcm: Long
    ): Pair<Int, Long> {
        inspections[monkey.name] = inspections[monkey.name]!! + 1

        var worry = monkey.operation(item)

        worry /= relief

        var throwTo = monkey.onFalse
        if (monkey.test(worry)) {
            throwTo = monkey.onTrue
        }

        worry %= lcm

        return Pair(throwTo, worry)
    }

    private fun Monkey.test(worry: Long): Boolean {
        return worry % divisible == 0L
    }


    private fun String.toOperation(): (worry: Long) -> Long {
        val (a, op, b) = """new\s=\s(old|\d+)\s([*+-/])\s(old|\d+)$""".toRegex().find(this)!!.destructured
        return when {
            a == "old" && b == "old" -> when (op) {
                "*" -> { worry -> worry * worry }
                "+" -> { worry -> worry + worry }
                "-" -> { _ -> 0 }
                else -> { _ -> 0 }
            }
            a == "old" && b != "old" -> when (op) {
                "*" -> { worry -> worry * b.toLong() }
                "+" -> { worry -> worry + b.toLong() }
                "-" -> { worry -> worry - b.toLong() }
                else -> { worry -> worry / b.toLong() }
            }
            else -> when (op) {
                "*" -> { worry -> a.toLong() * worry }
                "+" -> { worry -> a.toLong() + worry }
                "-" -> { worry -> a.toLong() - worry }
                else -> { worry -> a.toLong() / worry }
            }
        }
    }

    private fun initMonkeys(): List<Monkey> {
        return listOf(
            Monkey(
                name = 0,
                items = mutableListOf(54, 53),
                operation = "new = old * 3".toOperation(),
                divisible = 2,
                onTrue = 2,
                onFalse = 6
            ),
            Monkey(
                name = 1,
                items = mutableListOf(95, 88, 75, 81, 91, 67, 65, 84),
                operation = "new = old * 11".toOperation(),
                divisible = 7,
                onTrue = 3,
                onFalse = 4
            ),
            Monkey(
                name = 2,
                items = mutableListOf(76, 81, 50, 93, 96, 81, 83),
                operation = "new = old + 6".toOperation(),
                divisible = 3,
                onTrue = 5,
                onFalse = 1
            ),
            Monkey(
                name = 3,
                items = mutableListOf(83, 85, 85, 63),
                operation = "new = old + 4".toOperation(),
                divisible = 11,
                onTrue = 7,
                onFalse = 4
            ),
            Monkey(
                name = 4,
                items = mutableListOf(85, 52, 64),
                operation = "new = old + 8".toOperation(),
                divisible = 17,
                onTrue = 0,
                onFalse = 7
            ),
            Monkey(
                name = 5,
                items = mutableListOf(57),
                operation = "new = old + 2".toOperation(),
                divisible = 5,
                onTrue = 1,
                onFalse = 3
            ),
            Monkey(
                name = 6,
                items = mutableListOf(60, 95, 76, 66, 91),
                operation = "new = old * old".toOperation(),
                divisible = 13,
                onTrue = 2,
                onFalse = 5
            ),
            Monkey(
                name = 7,
                items = mutableListOf(65, 84, 76, 72, 79, 65),
                operation = "new = old + 5".toOperation(),
                divisible = 19,
                onTrue = 6,
                onFalse = 0
            )
        )
    }
}

data class Monkey(
    val name: Int,
    val items: MutableList<Long>,
    val operation: (worry: Long) -> Long,
    val divisible: Long,
    val onTrue: Int,
    val onFalse: Int
)

fun main() {
    val day = Day11()


    println("Solution 1: ${day.solve1()}")
    println("Solution 2: ${day.solve2()}")
}

