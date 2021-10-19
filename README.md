# ToDoList
ToDoList written in Golang

Bu projede ayrıca bir database kurulumu ile uğraşmamak adına SQLite kullandım.

Projede kullanılabilecek 3 request tipi var: 

Insert işlemi için :          POST: 127.0.0.1:8080/ToDoApp/v1/Insert

Update işlemi için:           PATCH: 127.0.0.1:8080/ToDoApp/v1/Update

Id'ye göre task bulmak için:  GET: 127.0.0.1:8080/ToDoApp/v1/GetByID

Başlangıç olarak database'de 4 adet Task bulunmakta. Task'lar "Title" bilgisine göre 'UNIQUE' olarak tanımlandığı için 
aynı title'a sahip birden fazla task oluşturulamayacak şekilde bir dizayn oluşturdum. Bu CreateDatabase.go dosyasında 
değişiklik yaparak kolayca değiştirilebilir. 
 


Yeni bir task oluşturup listeye eklemek için:

POST: 127.0.0.1:8080/ToDoApp/v1/Insert

    {
    
        "Title": "Ayşe'ye doğum günü hediyesi alınacak",
        
        "Description": "Acilll",
        
        "Category": "Doğum günleri",
        
        "Progress": "In progress",
        
        "Deadline": "2021-10-18T12:00:00+03:00"
        
    }
   
Bu durumda listeye task'in şu şekilde eklendiğini göreceksiniz:  

    {
    
        "Id": 5,
        
        "Title": "Ayşe'ye doğum günü hediyesi alınacak",
        
        "Description": "Acilll",
        
        "Category": "Doğum günleri",
        
        "Progress": "Overdue",
        
        "Deadline": "2021-10-18T12:00:00+03:00",
        
        "Priority": "Overdue",
        
        "CreatedTime": "2021-10-19T14:52:18.436933537+03:00",
        
        "UpdatedTime": "2021-10-19T14:52:18.436933537+03:00"
        
    }
    
Task'ların öncelik sıralaması şu şekilde:
1) Kullanıcı Progress: "Completed" olarak girerse Priority de "Completed" olarak tanımlanır.

2) Kullanıcı Progress: "Overdue" olarak girerse Priority de "Overdue" olarak tanımlanır.

3) Kullanıcı Progress'i Overdue veya Completed girmezse ve task'ın deadline geçtiyse Priority ve Progress "Overdue" olarak tanımlanır.

4) Kullanıcı Progress: "In progress" olarak girerse Priorty aşağıdaki değerleri alır:

    a) Task'ın Description'ında acil-Acil (case sensitive değil) ifadeleri geçiyorsa --> "Critical"
  
    b) Kullanıcı Priorty bilgisini girmediyse:
  
        a.1) Deadline'a kalan süre 0-4 saat aralığındaysa --> "Critical"
      
        a.2) Deadline'a kalan süre 4-24 saat aralığındaysa --> "Important"
      
        a.3) Deadline'a kalan süre 24-72 saat aralığındaysa --> "Normal"
      
        a.4) Deadline'a kalan süre 72 saatten fazlaysa --> "Low"
      
    c) Kullanıcı Priorty bilgisini girdiyse Priorty kullanıcının verdiği değer olarak tanımlanır.
  

 
 Var olan bir task'ın bir veya birden fazla field'ını update etmek için:
 
  PATCH: 127.0.0.1:8080/ToDoApp/v1/Update

    {
   
        "Title": "Ayşe'ye doğum günü hediyesi alınacak",
        
        "Description": "Hediye opsiyonları düşün.",
        
        "Category": "Doğum günleri",
        
        "Progress": "In progress",
        
        "Deadline": "2021-10-20T12:00:00+03:00"
        
    }
    
  Bu örnekte deadline'ı 20 Ekim'e çektim ve Description'ı acil barındırmayacak şekilde güncelledim. Sonuçta aşağıdaki verildiği üzere değiştirilen field'lar ile birlikte UpdatedTime da  güncellendi.
  
    {
   
        "Id": 5,
        
        "Title": "Ayşe'ye doğum günü hediyesi alınacak",
        
        "Description": "Hediye opsiyonları düşün",
        
        "Category": "Doğum günleri",
        
        "Progress": "In progress",
        
        "Deadline": "2021-10-20T12:00:00+03:00",
        
        "Priority": "Important",
        
        "CreatedTime": "2021-10-19T15:18:46.842749667+03:00",
        
        "UpdatedTime": "2021-10-19T15:34:55.87095553+03:00"
        
    }
  
Son olarak istendiği üzere Id'ye göre Task seçme işlemi aşağıdaki şekilde yapılabilir:

    GET: 127.0.0.1:8080/ToDoApp/v1/GetByID
    
    {
   
        "Id": 3
        
    }
    
Sonuç olarak Id'si 3 olan task getirilir:

    {
    
    "Id": 3,
    
    "Title": "ALES başvurusu",
    
    "Description": "Başvuru tarihi geçmeden önce, acil!.",
    
    "Category": "Yüksek Lisans",
    
    "Progress": "Overdue.",
    
    "Deadline": "2021-12-03T12:00:00+03:00",
    
    "Priority": "Overdue.",
    
    "CreatedTime": "2021-10-19T15:18:42.902663548+03:00",
    
    "UpdatedTime": "2021-10-19T15:18:42.902663548+03:00" 
        
    }
  
 
  
  
