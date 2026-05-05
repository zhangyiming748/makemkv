# makemkv

处理原盘中提取的 m2ts 文件为尽量保持原样的 mkv

## 功能特性

- 🎬 批量转换蓝光原盘中的 M2TS 文件为 MKV 格式
- 🔄 递归扫描指定目录及其所有子文件夹
- 🎵 支持选择是否将音频转换为 FLAC 格式
- 🔊 独立的 MKV 音频转 FLAC 工具
- 📝 视频和字幕流无损复制，保持原始质量
- 📋 自动日志记录，方便追踪转换过程

## 安装

### 方式一：下载预编译二进制包（推荐）

从 [GitHub Releases](https://github.com/zhangyiming748/makemkv/releases/latest) 下载最新版本的二进制文件。

#### 直接下载链接

| 平台 | 架构 | 下载链接 |
|------|------|----------|
| Windows | amd64 | [makemkv_windows_amd64](https://github.com/zhangyiming748/makemkv/releases/latest/download/makemkv_windows_amd64) |
| Windows | arm64 | [makemkv_windows_arm64](https://github.com/zhangyiming748/makemkv/releases/latest/download/makemkv_windows_arm64) |
| macOS | amd64 | [makemkv_darwin_amd64](https://github.com/zhangyiming748/makemkv/releases/latest/download/makemkv_darwin_amd64) |
| macOS | arm64 | [makemkv_darwin_arm64](https://github.com/zhangyiming748/makemkv/releases/latest/download/makemkv_darwin_arm64) |
| Linux | amd64 | [makemkv_linux_amd64](https://github.com/zhangyiming748/makemkv/releases/latest/download/makemkv_linux_amd64) |
| Linux | arm64 | [makemkv_linux_arm64](https://github.com/zhangyiming748/makemkv/releases/latest/download/makemkv_linux_arm64) |

### 方式二：从源码编译

```bash
go build -o makemkv.exe
```

## 使用方法

### 基本命令

```bash
# 查看帮助
./makemkv --help
./makemkv m2m --help
./makemkv flac --help

# 转换指定目录下的所有 M2TS 文件（不转换音频）
./makemkv m2m -d /path/to/blu-ray

# 转换并同时将音频转为 FLAC 格式
./makemkv m2m -d /path/to/blu-ray -f
./makemkv m2m -d /path/to/blu-ray --flac

# 将已有 MKV 文件的音频转换为 FLAC
./makemkv flac -d /path/to/mkv-folder
```

### 命令行参数

#### m2m 命令参数

| 参数 | 短格式 | 说明 | 默认值 |
|------|--------|------|--------|
| `--dir` | `-d` | 要搜索 M2TS 文件的根目录（必需） | - |
| `--flac` | `-f` | 是否将音频转换为 FLAC 格式 | false |
| `--keep` | `-k` | 转换后保留原始 M2TS 文件 | false |
| `--help` | `-h` | 显示帮助信息 | - |

#### flac 命令参数

| 参数 | 短格式 | 说明 | 默认值 |
|------|--------|------|--------|
| `--dir` | `-d` | 要搜索 MKV 文件的根目录（必需） | - |
| `--help` | `-h` | 显示帮助信息 | - |

### 示例

#### M2TS 转 MKV

```bash
# 示例 1: 转换目录，保持原始音频格式
./makemkv m2m -d "/vol2/1000/disk3/原盘/unzip/mount/BDMV/STREAM"

# 示例 2: 转换目录，音频转为 FLAC
./makemkv m2m -d "/vol2/1000/disk3/原盘/unzip/mount/BDMV/STREAM" -f

# 示例 3: 使用长参数
./makemkv m2m --dir "/path/to/folder" --flac=true

# 示例 4: 转换后保留原始 M2TS 文件
./makemkv m2m -d "/path/to/folder" -f -k
```

#### MKV 音频转 FLAC

```bash
# 示例 5: 将目录下所有 MKV 的音频转为 FLAC
./makemkv flac -d "/path/to/mkv-folder"

# 示例 6: 使用长参数
./makemkv flac --dir "/path/to/mkv-folder"
```

## 技术细节

### FFmpeg 转换参数

#### m2m 命令（M2TS → MKV）

程序使用以下 FFmpeg 参数进行转换：

```bash
ffmpeg -i input.m2ts \
  -c:v copy \           # 视频流无损复制
  -c:a flac \           # 音频转换为 FLAC（可选）
  -c:s copy \           # 字幕流无损复制
  -map 0 \              # 映射所有流
  -copy_unknown \       # 复制未知类型的流
  -fflags +genpts \     # 生成 PTS
  -avoid_negative_ts make_zero \  # 避免负时间戳
  -f matroska \         # 输出格式为 Matroska
  output.mkv
```

#### flac 命令（MKV 音频 → FLAC）

```bash
ffmpeg -i input.mkv \
  -c:v copy \           # 视频流无损复制
  -c:a flac \           # 音频转换为 FLAC
  -c:s copy \           # 字幕流无损复制
  output_tmp.mkv
```

### 日志系统

- 日志文件：`m2ts2mkv.log`
- 日志轮转：最大 1MB，保留 1 个备份，最多保存 28 天
- 时区：Asia/Shanghai
- 同时输出到控制台和日志文件

## 项目结构

```
makemkv/
├── core/
│   ├── mkv.go          # 核心转换逻辑
│   └── convert_test.go # 测试文件
├── main.go             # CLI 入口
├── go.mod              # Go 模块定义
└── README.md           # 项目说明
```

## 依赖

- [cobra](https://github.com/spf13/cobra) - CLI 框架
- [finder](https://github.com/zhangyiming748/finder) - 文件查找工具
- [lumberjack](https://github.com/zhangyiming748/lumberjack) - 日志轮转
- FFmpeg - 视频转换工具（需要预先安装在系统中）

## 注意事项

1. 确保系统中已安装 FFmpeg 并添加到 PATH
2. 转换后的 MKV 文件会保存在与源文件相同的目录
3. 单个文件转换失败不会影响其他文件的处理
4. 建议先在小范围测试，确认效果后再批量处理

## License

详见 [LICENSE](LICENSE) 文件
