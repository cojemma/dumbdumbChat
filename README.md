# dumbdumbChat
dumbdumbChat是一個使用chatGPT作為語言模型,搭配azure提供的語音生成技術以及live2d模型驅動的虛擬youtuber。
## 安裝方法
AIVtuber需要以下的環境與API: 
- OpenAI API Key: 至少一個OpenAI的API Key以使用GPT-3模型。  
<https://platform.openai.com/account/api-keys> 
- Azure 語音生成資源: 如果需要語音生成功能需要在Azure上申請語音合成的訂閱。
<https://speech.microsoft.com/portal/voicegallery>
## 使用方法  
1. 啟動aiVtuber.exe後,網頁服務將在`http://localhost:8001/` 
2. 於網頁上填入OpenAI API Key以及Azure語音訂閱的相關資訊。  
3. 即可開始與AIVtuber進行語音對話! 
4. 如果要修改給ChatGPT的角色設定與對話，位於.chatAI/chathistory.json。
5. 如果要新增live2d的話，將含有模型的model(model3).json檔案的資料夾放入./static/live2d的資料夾內。之後重啟執行檔便可以讀取模型選項。
6. 可以於Set Live2d Config下設定live2d模型對應語句情緒所表現的Motion及Expression，輸入的動作需於live2d中已包含的動作。
## 檔案結構
chatAI - chatGpt request, chat history saved   
live2Drive - use emotion struct to drive live2D model   
tts - azure tts request.
## TODO
1. ~~模型motion及expression接口開放~~
2. ~~tts聲音模型新增可選選項~~
3. 更便宜的文字生成模型服務,e.x. Claude    
4. 更靈活的對話歷史及prompt編輯
## 聯絡資訊
如果有任何問題歡迎聯絡我:
- 信箱:iam880401@gmail.com
