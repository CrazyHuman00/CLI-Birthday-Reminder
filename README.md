# CLI-Birthday-Reminder

---

友達や身内の誕生日を管理する`birthday`コマンドです。[This page is English version](#cli-birthday-reminder-1)

- [概要](#概要)
- [DBに保存する内容](#dbに保存する内容)
- [コマンドの各機能](#コマンドの各機能)
    - [バージョン確認](#バージョン確認)
    - [誕生日リストの確認](#誕生日リストの確認)
    - [誕生日の追加](#誕生日の追加)
    - [誕生日の更新](#誕生日の更新)
    - [誕生日の削除](#誕生日の削除)
- [使用技術](#使用技術)
- [今後の方針](#今後の方針)
- [参考文献](#参考文献)

---

## 概要

CLIツールの起動時に誰の誕生日が近いかを表示してくれるコマンドです。CLI上でもDBに保存されている誕生日を追加、更新、削除することができます。誕生日の10日前から報告し、誕生日の10日後まで報告をする機能となっています。

## DBに保存する内容

- UserBirthday
    - id(PK): int
    - name: string
    - day: string

## コマンドの各機能

### バージョン確認

コマンドのバージョンを確認します。
```
birthday version
```

### 誕生日リストの確認

保存されている誕生日のリストを表示します。
```
birthday list
```


### 誕生日の追加

誕生日をユーザ名と日付とセットで保存します。
```
birthday add <username> <XX/XX>
```

### 誕生日の更新

ユーザ名から保存されている誕生日を更新します。
```
birthday update <username> <XX/XX>
```

### 誕生日の削除

ユーザ名から保存されている誕生日を削除します。
```
birthday remove <username>
```

## 使用技術

- Go
- [github.com/spf13/cobra](https://github.com/spf13/cobra)

## 今後の方針

- 人によって色をつけれるようにする。

## 参考文献

[GoでCLIコマンドを自作して公開するまでの道のり](https://zenn.dev/tttol/articles/c7dfc74d27e45d)

---

# CLI-Birthday-Reminder

The `birthday` command manages birthdays of friends and relatives.

- [Summary](#summary)
- [Contents to be saved to DB](#contents-to-be-saved-to-db)
- [Feature](#feature)
    - [version](#version)
    - [list](#list)
    - [add](#add)
    - [update](#update)
    - [remove](#remove)
- [Stack](#stack)

## Summary

This command displays who's birthday is coming up when you start the CLI tool. you can also add, update, or delete birthdays stored in the DB on the CLI. The function reports from 10 days before the birthday and up to 10 days after the birthday.

## Contents to be saved to DB

- UserBirthday
    - id(PK): int
    - name: string
    - day: string

## Feature

### version

Check the version of the command.
```
birthday version
```

### list

Displays a list of stored birthdays.
```
birthday list
```

### add

Save the birthday as a set with the user name and date.
```
birthday add <username> <XX/XX>
```

### update

Updates the stored birthdays from the user name.
```
birthday update <username> <XX/XX>
```

### remove

Delete stored birthdays from the user name.
```
birthday remove <username>
```

## Stack

- Golang
- [github.com/spf13/cobra](https://github.com/spf13/cobra)