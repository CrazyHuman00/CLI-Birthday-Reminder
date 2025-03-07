# CLI-Birthday-Reminder

- [概要](#概要)
- [DBに保存する内容](#dbに保存する内容)
- [コマンドの各機能](#コマンドの各機能)
    - [バージョン確認](#バージョン確認)
    - [誕生日リストの確認](#誕生日リストの確認)
    - [誕生日の追加](#誕生日の追加)
    - [誕生日の更新](#誕生日の更新)
    - [誕生日の削除](#誕生日の削除)
- [使用技術](#使用技術)

---

## 概要

CLIツールの起動時に誰の誕生日が近いかを表示してくれるコマンドです。
CLI上でもDBに保存されている誕生日を追加、更新、削除することができます。
保存されている誕生日と一致した日に起動するとその人が誕生日であると報告してくれます。
誕生日の10日前から報告し、誕生日の10日後まで報告をする機能となっています。

## DBに保存する内容

- UserBirthday
    - id(PK): int
    - name: string
    - day: string

## コマンドの各機能

### バージョン確認

コマンドのバージョンを確認します。
```
birthday-reminder version
```

### 誕生日リストの確認

保存されている誕生日のリストを表示します。
```
birthday-reminder list
```


### 誕生日の追加

誕生日をユーザ名と日付とセットで保存します。
```
birthday-reminder add <username> <XX/XX>
```

### 誕生日の更新

ユーザ名から保存されている誕生日を更新します。
```
birthday-reminder update <username> <XX/XX>
```

### 誕生日の削除

ユーザ名から保存されている誕生日を削除します。
```
birthday-reminder remove <username>
```

## 使用技術

- Go
- github.com/spf13/cobra

## 参考文献

---