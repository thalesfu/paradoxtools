# paradoxtools

Paradox 游戏文件格式解析库。

## 子模块

- `CK2/save/`：CK2 存档文件（`.ck2`）解析器，输出 `SaveFile{Characters, Dynasties, Provinces, Titles, ...}`
- `CK2/static/`：CK2 静态数据（特性、文化、宗教、修正、建筑、本地化文本）解析
- `utils/pserialize/`：Paradox 文本格式（类 Clausewitz 语法）序列化 / 反序列化

## 使用

```go
import "github.com/thalesfu/paradoxtools/CK2/save"

save, err := save.ParseSaveFile(path)
// save.Characters, save.Dynasties, save.Titles, save.Provinces ...
```
