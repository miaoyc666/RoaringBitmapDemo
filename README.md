# RoaringBitmap Demo

## 📌 项目简介
本项目演示了如何在 **Go 语言** 中使用 [RoaringBitmap](https://github.com/RoaringBitmap/roaring) 来存储和查询大规模 UID 集合。  
主要场景：
- 每个分类有几千万个 UID
- 需要快速获取某分类的 UID 集合
- 需要对多个分类进行交集运算（例如推荐、标签筛选等场景）

## 🔧 功能模块
- **prepare.go**
    - 用于准备数据，将每个分类的 UID 存储为 `.roar` 文件
    - 底层使用 RoaringBitmap 的高效压缩存储格式
- **main.go**
    - 用于加载 `.roar` 文件
    - 演示集合的基本统计（基数 cardinality）
    - 演示集合之间的交集查询

## 📂 文件结构
```bash
.
├── prepare.go # 数据生成（写入 .roar 文件）
├── main.go # 加载数据并执行交集查询
├── sports.roar # 示例分类数据（运行 prepare.go 后生成）
├── music.roar
├── movie.roar
└── README.md # 使用说明
```

## 🚀 使用步骤

### 1. 克隆代码并安装依赖
```bash
git clone <your-repo-url>
cd <repo-dir>

go mod init example.com/roaringdemo
go get github.com/RoaringBitmap/roaring/roaring64
```

### 2. 生成示例数据
运行：
```bash
go run prepare.go
```
输出示例：
```bash
Saved sports.roar with 5 UIDs
Saved music.roar with 4 UIDs
Saved movie.roar with 5 UIDs
```
生成的文件：
```bash
sports.roar
music.roar
movie.roar
```
### 3. 加载数据并查询
运行：
```bash
go run main.go
```
输出示例：
```bash
sports size: 5
music size: 4
movie size: 5
sports ∩ music cardinality: 2
common UID: 2
common UID: 3
```

## 📊 关键点说明
- .roar 文件是跨语言兼容的二进制格式（同一库在 Python、Java、Go 中都能读写）
- And() 是就地修改操作，会改变调用者的集合。若要保留原集合，请先 Clone()
- 常见操作：
  - 并集：bm1.Or(bm2)
  - 交集：bm1.And(bm2)
  - 差集：bm1.AndNot(bm2)
  - 基数统计：bm1.GetCardinality()

## 📝 应用场景
- 用户标签交集计算（推荐、广告定向）
- 大规模去重 / 过滤
- 高效集合统计（比传统 set 占用更少内存）
