package main

import (
	"fmt"
	"slices"
)

// https://takeuforward.org/data-structure/minimum-number-of-platforms-required-for-a-railway/

// Optimal approach
func findPlatform(arr []int, dep []int) int {
	if len(arr) <= 1 {
		return 1
	}

	slices.Sort(arr)
	slices.Sort(dep)

	left := 0
	right := 1
	maxCount := 1
	count := 1
	mcount := len(arr)

	for left < mcount && right < mcount {
		if arr[right] <= dep[left] {
			count++
			right++
		} else {
			count--
			left++
		}

		maxCount = max(maxCount, count)
	}

	return maxCount
}

// Just sorting the departure.
// func findPlatform(arr []int, dep []int) int {
// 	if len(arr) <= 1 {
// 		return 1
// 	}

// 	mcount := len(arr)
// 	meetings := []int{}
// 	for i := 0; i < mcount; i++ {
// 		meetings = append(meetings, i)
// 	}

// 	slices.SortFunc(meetings, func(i, k int) int {
// 		if dep[i] < dep[k] {
// 			return -1
// 		} else if dep[i] > dep[k] {
// 			return 1
// 		} else {
// 			return 0
// 		}
// 	})

// 	left := 0
// 	right := 1
// 	maxCount := 1

// 	for left < mcount-1 {
// 		counter := 1

// 		right = left + 1

// 		for right < mcount {
// 			if arr[meetings[right]] <= dep[meetings[left]] {
// 				counter++
// 			}

// 			right++
// 		}

// 		maxCount = max(maxCount, count)

// 		left++
// 	}

// 	return maxCount
// }

func main() {
	fmt.Println(findPlatform(
		[]int{900, 945, 955, 1100, 1500, 1800},
		[]int{920, 1200, 1130, 1150, 1900, 2000},
	))

	fmt.Println(findPlatform(
		[]int{900, 945, 950, 1100, 1500, 1800},
		[]int{910, 1200, 1120, 1130, 1900, 2000},
	))

	fmt.Println(findPlatform(
		[]int{900, 945, 950, 1100, 1500, 1800, 2200},
		[]int{910, 1200, 1120, 1130, 1900, 2000, 2300},
	))

	fmt.Println(findPlatform(
		[]int{1026, 445, 145, 555, 710, 1712, 1105, 506, 531, 1930, 220, 611, 553, 53, 401, 2000, 225, 1359, 1159, 120, 1857, 740, 253, 303, 740, 1434, 1407, 807, 423, 500, 851, 809, 527, 123, 1117, 23, 1050, 1613, 1025, 1225, 826, 422, 221, 28, 515, 401, 1241, 38, 315, 1007, 508, 1054, 803, 333, 11, 225, 137, 30, 344, 36, 841, 1419, 552, 422, 337, 1222, 1422, 10, 258, 1434, 538, 153, 808, 347, 104, 1136, 302, 357, 938, 2015, 403, 258, 736, 1057, 1547, 531, 1642, 2333, 511, 1301, 1059, 638, 117, 314, 905, 1314, 1103, 356, 6, 235, 1209, 849, 304, 539, 921, 625, 741, 47, 55, 127, 3, 53, 608, 1725, 1431, 1427, 814, 935, 1435, 938, 448, 930, 549, 1002, 909, 329, 1203, 824, 212, 852, 1825, 541, 1136, 825, 110, 1301, 1619, 31, 37, 140, 443, 831, 1220, 346, 1242, 1414, 1702, 113, 901, 1221, 742, 941, 711, 825, 220, 1526, 125, 6, 1418, 1632, 1020, 1957, 711, 552, 430, 157, 359, 212, 1150, 2212, 857, 1351, 1531, 219, 6, 813, 141, 1730, 458, 911, 1201, 218, 441, 1121, 216, 501, 528, 1049, 158, 646, 1303, 18, 123, 1408, 1750, 604, 1042, 1, 1049, 343, 821, 154, 139, 1138, 1138, 2107, 318, 556, 1259, 44, 457, 116, 1555, 1546, 525, 1648, 1612, 504, 1218, 2231, 1408, 922, 1253, 353, 759, 727, 817, 15, 650, 519, 1723, 1329, 357, 839, 440, 824, 805, 1524, 943, 116, 937, 740, 357, 1023, 314, 931, 1844, 1639, 232, 23, 549, 28, 1135, 1220, 449, 925, 1112, 1652, 609, 2316, 1205, 133, 223, 543, 1004, 1426, 814, 422, 1134, 228, 832, 641, 1643, 1339, 1725, 1023, 1017, 1642, 627, 642, 934, 1349, 624, 1411, 246, 508, 1113, 955, 14, 258, 502, 749, 1423, 400, 809, 157, 2100, 120, 935, 1417, 1204, 1016, 144, 724, 1221, 207, 1745, 1905, 333, 324, 1407, 1537, 1246, 1543, 226, 1031, 1045, 1325, 957, 419, 528, 419, 1239, 131, 452, 834, 3, 430, 544, 348, 538, 952, 1956, 24, 1140, 407, 800, 26, 1101, 1301, 838, 555, 816, 130, 1645, 25, 13, 136, 810, 1624, 819, 811, 1138, 622, 1643, 1113, 1434, 1311, 1501, 847, 950, 356, 305, 1825, 1135, 946, 1550, 211, 835, 539, 335, 1123, 904, 938, 812, 1339, 1921, 1511, 709, 130, 1222, 820, 405, 330, 1254, 432, 130, 33, 254, 826, 343, 830, 1708, 657, 1149, 1148, 406, 1204, 1253, 1006, 421, 258, 2022, 1027, 110, 1426, 458, 128, 39, 537, 737, 1938, 849, 424, 1332, 140, 1123, 103, 250, 600, 749, 247, 1243, 108, 138, 1119, 443, 409, 343, 654, 413, 658, 333, 22, 739, 453, 51, 104, 240, 1201, 1328, 208, 242, 1020, 1305, 708, 630, 1314, 1018, 1104, 2137, 224, 1051, 543, 838, 539, 606, 949, 1435, 1034, 402, 1105, 936, 530, 946, 1629, 1641, 744, 23, 508, 36, 300, 756, 302, 733, 600, 2009, 804, 653, 1312, 24, 1217, 345, 301, 1450, 103, 1512, 905, 0, 12, 854, 451, 317, 336, 125, 400, 829, 202, 927, 257, 511, 406, 1548, 453, 1518, 1057, 2016, 522, 938, 346, 536, 1723, 531, 727, 232, 1050, 155, 738, 2012, 957, 555, 504, 652, 1301, 522, 351, 136, 511, 334, 410, 933, 805, 1603, 1808, 105, 325, 1814, 24, 552, 401, 41, 1412, 800, 133, 7, 755, 2005, 1258, 1034, 1215, 14, 1553, 807, 310, 1416, 246, 243, 26, 857, 2157, 1605, 555, 1322, 253, 1130, 702, 1350, 50, 105, 345, 121, 350, 300, 1616, 448, 1008, 737, 629, 700, 945, 45, 817, 1947, 549, 1451, 1732, 935, 135, 132, 408, 1407, 1801, 1640, 50, 313, 21, 308, 1320, 1207, 2216, 1918, 111, 348, 621, 30, 842, 818, 28, 1056, 1941, 1635, 1111, 508, 420, 139, 1036, 717, 1250, 320, 536, 18, 156, 808, 1754, 451, 432, 236, 1128, 912, 1326, 1525, 526, 355, 439, 223, 1702, 2109, 503, 1629, 224, 107, 1649, 1338, 722, 840, 923, 1, 144, 102, 1128, 839, 329, 443, 1402, 1132, 226, 422, 456, 157, 1126, 458, 1114, 600, 12, 301, 313, 1213, 210, 1110, 205, 34, 42, 320, 313, 101, 253, 301, 311, 1539, 1911, 835, 1556, 1049, 17, 530, 1604, 2, 410, 117, 428, 1638, 930, 349, 708, 1327, 118, 929, 402, 437, 2143, 424, 751, 230, 2321, 30, 1555, 1446, 404, 1135, 914, 1305, 215, 455, 356, 1216, 302, 815, 1357, 952, 1051, 36, 59, 30, 147, 403, 1720, 722, 1509, 423, 207, 428, 214, 1227, 448, 111, 206, 248, 313, 246, 424, 1416, 738, 1240, 558, 1625, 548, 739, 134, 704, 656, 836, 302, 553, 105, 101, 115, 1438, 208, 244, 1503, 1231, 1619, 616, 934, 1524, 800, 740, 1146, 347, 1002, 1550, 943, 657, 256, 559, 234, 745, 1002, 709, 801, 345, 1138, 536, 2214, 659, 427, 850, 145, 1938, 726, 328, 407, 1512, 358, 439, 538, 326, 2305, 15, 854, 1636, 403, 1556, 218, 1139, 1319, 1042, 216, 1103, 206, 1519, 920, 228, 203, 444, 911},
		[]int{1713, 2242, 1144, 848, 1941, 1734, 2347, 1726, 2247, 2018, 355, 2249, 2134, 758, 2044, 2354, 237, 2152, 1221, 532, 2031, 820, 1013, 2311, 2150, 2321, 1909, 2344, 1732, 2127, 2126, 1602, 945, 705, 1632, 1305, 1604, 1639, 1630, 1334, 1858, 2131, 350, 1625, 1443, 926, 1245, 1802, 1558, 1406, 1442, 2024, 1450, 703, 508, 1341, 1445, 1624, 1414, 2143, 1306, 1808, 845, 1717, 1928, 1620, 1631, 1101, 2146, 2153, 1524, 2306, 944, 702, 1219, 1318, 1431, 2044, 1616, 2106, 1631, 841, 1128, 1732, 1957, 719, 2052, 2342, 2016, 1348, 1505, 957, 1749, 509, 2140, 1515, 2114, 1245, 1806, 1439, 1951, 2341, 518, 2315, 2304, 1117, 1701, 1837, 346, 2200, 118, 1537, 714, 1757, 2237, 2139, 2103, 1707, 2031, 1848, 1623, 2105, 2125, 1205, 1730, 512, 1928, 1816, 2250, 1624, 1828, 1227, 2308, 1044, 809, 1636, 2101, 321, 121, 1237, 1157, 1633, 1731, 2050, 1419, 1724, 2155, 705, 2109, 2343, 844, 2246, 2037, 1452, 1131, 2135, 1441, 549, 1810, 2039, 2306, 2215, 1057, 739, 2323, 1250, 1024, 1856, 2152, 2238, 2323, 1855, 1920, 454, 1008, 2059, 1556, 2202, 1645, 2320, 1843, 1812, 1011, 2052, 1639, 1555, 2234, 2025, 1502, 2224, 1742, 957, 1527, 2110, 1929, 1712, 2159, 239, 2113, 1841, 2211, 1221, 1028, 2124, 1614, 2139, 2240, 926, 1527, 1839, 721, 2316, 1931, 1648, 817, 1707, 2002, 921, 2342, 2336, 1842, 2257, 1925, 827, 1800, 1349, 1949, 1242, 906, 1954, 1732, 1721, 2010, 2349, 2034, 1527, 1328, 2000, 2130, 1804, 1745, 1105, 1406, 2245, 1649, 1411, 2300, 2330, 713, 553, 2329, 632, 1556, 1942, 1446, 1510, 1253, 2341, 1532, 2322, 1810, 445, 1953, 1619, 1436, 2145, 1115, 715, 1245, 2126, 1834, 2101, 2326, 1850, 1736, 1346, 1914, 1806, 727, 2249, 1611, 1700, 1410, 1910, 1831, 1355, 2136, 2157, 229, 446, 1022, 1914, 1709, 1717, 1442, 1748, 2156, 651, 1424, 1951, 2304, 1154, 2200, 1059, 1510, 2047, 2142, 2013, 1303, 901, 2134, 1927, 1256, 2205, 709, 2129, 1602, 1336, 1045, 821, 2352, 1559, 1626, 1224, 627, 2134, 1348, 1827, 1508, 1001, 1051, 1546, 2101, 1934, 2011, 2339, 1807, 953, 2005, 1542, 2144, 1448, 2203, 452, 1934, 1309, 622, 853, 1814, 1812, 1557, 2039, 1920, 1627, 2149, 1955, 1519, 1312, 1702, 2252, 1107, 1049, 420, 2128, 1717, 2146, 1751, 1413, 2335, 1815, 817, 1907, 2211, 1023, 2318, 1423, 1957, 1616, 2012, 2308, 1259, 2114, 808, 1527, 2135, 1638, 1956, 408, 655, 1346, 527, 1248, 2055, 2235, 1358, 1501, 907, 2101, 2244, 2108, 929, 2200, 2056, 1532, 256, 1900, 2218, 1358, 214, 1414, 1105, 2311, 1232, 611, 1848, 1543, 1459, 743, 1143, 1905, 2241, 1758, 1446, 136, 2323, 1705, 2057, 504, 2029, 1708, 1642, 1945, 926, 802, 1821, 1028, 1145, 1527, 1531, 1419, 1730, 243, 324, 2131, 2117, 2022, 1426, 1836, 1540, 1319, 2337, 2304, 1948, 1607, 1220, 1407, 1804, 2209, 2025, 1412, 1729, 2134, 1353, 905, 1529, 1647, 2345, 848, 835, 947, 1710, 1325, 1519, 2255, 803, 1008, 2048, 2005, 1952, 2328, 1717, 1523, 2043, 2141, 1711, 1030, 2211, 1522, 202, 43, 1022, 1848, 1457, 1313, 448, 722, 1601, 2005, 1042, 1802, 2114, 1145, 2253, 2209, 2210, 1705, 2128, 624, 1345, 1648, 2009, 2242, 1248, 2214, 1019, 2020, 1447, 2144, 2119, 2056, 2356, 1646, 1002, 2331, 657, 428, 422, 608, 1149, 2047, 1900, 1241, 2358, 2201, 755, 1351, 2317, 1623, 2205, 1518, 2210, 1518, 1657, 1838, 2026, 822, 2203, 1449, 1639, 1313, 2241, 1814, 1023, 1611, 1626, 708, 1743, 125, 2302, 2254, 2143, 1922, 1710, 604, 1204, 2020, 1414, 1947, 409, 1747, 626, 751, 1729, 1710, 2334, 2304, 1750, 839, 1120, 1929, 1640, 1605, 1951, 1433, 1607, 2306, 1343, 2203, 622, 2341, 1705, 1905, 2318, 1218, 1946, 1354, 344, 1803, 1858, 2229, 2237, 2303, 946, 1420, 1956, 1118, 2246, 1957, 1738, 2038, 2217, 1217, 1236, 836, 1816, 1744, 832, 1932, 2300, 2037, 1304, 514, 1004, 2257, 1722, 2112, 1221, 2035, 1956, 1918, 1833, 2225, 1823, 1505, 1952, 2343, 2110, 1611, 2024, 233, 1220, 1935, 1750, 1457, 1212, 1003, 1126, 1034, 454, 1922, 1714, 2323, 1619, 1446, 1943, 1543, 2243, 1114, 840, 2212, 2316, 1258, 1126, 629, 514, 1416, 1608, 1821, 1442, 842, 1554, 1251, 1231, 1030, 1804, 2051, 2339, 714, 2109, 2232, 1423, 2351, 1853, 1919, 1637, 2250, 1155, 2106, 2029, 1646, 2330, 1719, 2153, 1156, 1925, 417, 2039, 1118, 2131, 2238, 2355, 1730, 2142, 2357, 1210, 1956, 1626, 611, 1409, 2256, 2243, 1157, 2324, 1509, 2001, 1022, 1845, 2333, 2255, 1305, 935, 646, 113, 1348, 2333, 2304, 1522, 2207, 2018, 408, 1657, 1757, 1547, 1404, 849, 1549, 2102, 331, 1622, 1827, 2135, 2307, 1635, 1137, 1805, 945, 1822, 2027, 1657, 954, 2358, 340, 850, 1720, 811, 715, 2121, 1826, 1626, 2004, 1405, 1931, 2246, 2047, 2218, 1022, 1439, 1401, 1900, 2130, 2328, 2119, 1851, 428, 1235, 1811, 1040, 1906, 2212, 1937, 1405, 2210, 710, 2359, 1021, 759, 1406, 312, 2110, 1238, 839, 2110, 2021, 854, 850, 912, 2329, 2350, 1253, 1802, 2114, 632, 2350, 900, 2204, 1853, 1419, 1038, 1556, 259, 1601, 1738, 1701, 2302, 2140, 1147},
	))

}