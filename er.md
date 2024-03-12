# データベース設計
## テーブル
### - Users
全体のアカウント情報を管理するテーブル

|Field|Type|Require|About|
|-|-|-|-|
|id|UUID|True|ユーザーID|
|type|enum|True|アカウントタイプ [Mobile / Card]|
|email|String|False|メールアドレス|
|password_hash|String|False|パスワード|
|created_at|Timestamp|False|作成日時|
|updated_at|Timestamp|False|更新日時|

### - Community
コミュニティー情報を管理するテーブル
(現状は名前と詳細の管理のみを行うが今後は画像設定などもできると良いかも)

|Field|Type|Require|About|
|-|-|-|-|
|id|Int|True|コミュニティーID|
|name|String|True|コミュニティー名|
|discription|String|False|コミュニティー説明|
|rate|Dboule|True|基本通貨との交換レート <br> (基本通貨=コミュニティー通貨*rate)|
|tax|Double|True|送金手数料|
|owner_user_id|UUID|False|コミュニティーオーナーのユーザーID|
|created_at|Timestamp|False|作成日時|
|updated_at|Timestamp|False|更新日時|

### - Wallet
全てのウォレット情報を管理するテーブル

|Field|Type|Require|About|
|-|-|-|-|
|id|Int|True|ウォレットID|
|balance|Int|True|残高|
|user_id|UUID|False|所有者のユーザーID|
|community_id|Int|False|所属コミュニティーID|
|created_at|Timestamp|False|作成日時|
|updated_at|Timestamp|False|更新日時|

### - Transaction
流通情報を管理するテーブル

|Field|Type|Require|About|
|-|-|-|-|
|id|UUID|True|トランザクションID|
|amount|Int|True|送金量|
|tax|Int|True|手数料|
|total_amount|Int|True|総額|
|transaction_date|Timestamp|True|実行日時|
|from_community_id|Int|False|送金元コミュニティーID|
|to_community_id|Int|False|送金先コミュニティーID|
|from_wallet_id|Int|False|送金元ウォレットID|
|to_wallet_id|Int|False|送金先ウォレットID|
|created_at|Timestamp|False|作成日時|
|updated_at|Timestamp|False|更新日時|