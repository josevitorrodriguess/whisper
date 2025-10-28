# 📱 Whisper - Rede Social de Postagens

## 🎯 Visão Geral
Uma rede social minimalista focada em postagens de texto, onde usuários podem compartilhar pensamentos, interagir através de curtidas e comentários, e receber notificações em tempo real.

---

## 🛠️ Stack Tecnológica

### Backend
- **Linguagem:** Go (Golang)
- **ORM:** GORM
- **Framework Web:** Gin ou Fiber (recomendado Gin pela maturidade)
- **Autenticação:** Firebase Admin SDK para validar tokens
- **Banco de Dados:** PostgreSQL

### Frontend
- **Framework:** Next.js 14+ (App Router)
- **Autenticação:** Firebase Authentication
- **Estilização:** Tailwind CSS
- **Estado:** React Context ou Zustand
- **Requisições:** Axios ou Fetch API

### Infraestrutura e Deploy

#### ⚠️ IMPORTANTE - Entendendo o Deploy:

**Frontend (Next.js):**
- **Plataforma:** Vercel
- **Certeza:** 95% ✅
- **Como funciona:** 
  - Você conecta seu repositório GitHub à Vercel
  - Cada push para main faz deploy automático
  - Vercel é feita pela mesma empresa do Next.js, integração perfeita
  - Domínio gratuito .vercel.app

**Backend (Go):**
- **Plataforma:** Railway
- **Certeza:** 85% ⚠️
- **Como funciona:**
  - Railway detecta automaticamente que é Go
  - Cria um Dockerfile se não existir
  - Gera uma URL pública para sua API
  - Você precisa de um Procfile ou configurar o comando de start
- **Alternativas:** Render.com (mais fácil), Fly.io (mais controle)
- **Custo:** Railway tem plano grátis limitado ($5 de crédito/mês)

**Banco de Dados (PostgreSQL):**
- **Opção 1 - Supabase (Recomendado):**
  - **Certeza:** 90% ✅
  - **Prós:** 
    - Interface visual para gerenciar dados
    - Backup automático
    - 500MB grátis
    - Fornece URL de conexão pronta
  - **Contras:** 
    - Banco "dorme" se não usado (plano grátis)
    
- **Opção 2 - Railway PostgreSQL:**
  - **Certeza:** 85% ⚠️
  - **Prós:** 
    - Tudo no mesmo lugar (backend + DB)
    - Fácil conectar backend ao DB (rede interna)
  - **Contras:** 
    - Consome do mesmo crédito grátis do backend
    
- **Opção 3 - Render PostgreSQL:**
  - **Certeza:** 80% ⚠️
  - **Prós:** Gratuito permanente
  - **Contras:** Expira após 90 dias no plano grátis

**⚠️ AVISO:** Nunca fiz deploy de banco separado do backend antes. A configuração mais segura para começar seria:
- **Supabase** para PostgreSQL (mais confiável, tem UI)
- **Railway** para o backend Go (mais simples que Render)
- **Vercel** para Next.js (obviamente)

**Como conectar tudo:**
1. Supabase te dá uma DATABASE_URL (exemplo: `postgresql://user:pass@db.supabase.co:5432/postgres`)
2. Você coloca essa URL como variável de ambiente no Railway
3. Seu backend Go conecta usando GORM
4. Frontend chama a API do Railway (você precisa habilitar CORS no Go)

---

## 📐 Arquitetura do Sistema

```
┌─────────────────┐
│   Next.js App   │
│   (Vercel)      │
└────────┬────────┘
         │
         │ HTTP/REST
         │
    ┌────▼────────────────┐
    │   Go Backend API    │
    │   (Railway)         │
    │                     │
    │  ┌──────────────┐   │
    │  │ Firebase SDK │   │  (Valida tokens)
    │  └──────────────┘   │
    └────────┬────────────┘
             │
             │ SQL
             │
    ┌────────▼────────────┐
    │   PostgreSQL        │
    │   (Supabase)        │
    └─────────────────────┘
```

---

## 🔥 Funcionalidades Detalhadas

### 1. Autenticação (Firebase)
- [x] Cadastro de usuário (email/senha)
- [x] Login
- [x] Logout
- [x] Recuperação de senha
- [x] Profile do usuário (nome, avatar)

**Fluxo:**
1. Frontend usa Firebase Auth para login
2. Firebase retorna um token JWT
3. Frontend envia token no header de todas requisições
4. Backend valida token usando Firebase Admin SDK

### 2. Posts

#### Criar Post
- **Limite de caracteres:** 500 caracteres (você pode ajustar)
- **Campos:**
  - Conteúdo (texto)
  - Data/hora de criação (auto)
  - Autor (ID do usuário)

#### Editar Post
- Apenas o autor pode editar
- Marcar como "editado" com timestamp
- Histórico de edições? (opcional)

#### Deletar Post
- Apenas o autor pode deletar
- Deletar em cascata comentários e curtidas

#### Listar Posts
- Feed principal (todos os posts, ordem cronológica reversa)
- Posts de um usuário específico
- Paginação (10-20 posts por página)

### 3. Curtidas
- Um usuário pode curtir um post apenas uma vez
- Descurtir disponível
- Contador de curtidas visível
- **Notificação:** Quando alguém curte, autor do post recebe notificação

### 4. Comentários
- Comentar em posts
- Editar comentário (apenas autor)
- Deletar comentário (autor ou dono do post)
- Listar comentários de um post
- Contador de comentários

**⚠️ DÚVIDA (70% certeza):** Comentários podem ter curtidas também? Deixei fora por enquanto.

### 5. Notificações
- Tipos:
  - Alguém curtiu seu post
  - Alguém comentou no seu post
  - ⚠️ (Opcional) Alguém respondeu seu comentário
- Status: lida/não lida
- **Implementação:** Polling (frontend consulta a cada X segundos) ou WebSocket

**⚠️ AVISO (60% certeza):** WebSocket em Go + Next.js + Railway pode ser complexo. Recomendo começar com polling simples.

---

## 🗄️ Modelo de Dados (PostgreSQL + GORM)

### User (gerenciado pelo Firebase, mas precisamos de uma tabela local)
```go
type User struct {
    ID          string    `gorm:"primaryKey"` // Firebase UID
    Email       string    `gorm:"unique;not null"`
    DisplayName string    `gorm:"not null"`
    PhotoURL    string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Post
```go
type Post struct {
    ID         uint      `gorm:"primaryKey"`
    UserID     string    `gorm:"not null;index"` // FK para User
    Content    string    `gorm:"type:text;not null"`
    IsEdited   bool      `gorm:"default:false"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
    
    // Relacionamentos
    User       User        `gorm:"foreignKey:UserID"`
    Likes      []Like      `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
    Comments   []Comment   `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}
```

### Like
```go
type Like struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    string    `gorm:"not null;index"`
    PostID    uint      `gorm:"not null;index"`
    CreatedAt time.Time
    
    // Relacionamentos
    User User `gorm:"foreignKey:UserID"`
    Post Post `gorm:"foreignKey:PostID"`
}

// Índice único para garantir um like por usuário/post
// gorm:"uniqueIndex:idx_user_post"
```

### Comment
```go
type Comment struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    string    `gorm:"not null;index"`
    PostID    uint      `gorm:"not null;index"`
    Content   string    `gorm:"type:text;not null"`
    IsEdited  bool      `gorm:"default:false"`
    CreatedAt time.Time
    UpdatedAt time.Time
    
    // Relacionamentos
    User User `gorm:"foreignKey:UserID"`
    Post Post `gorm:"foreignKey:PostID"`
}
```

### Notification
```go
type Notification struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      string    `gorm:"not null;index"` // Quem recebe
    ActorID     string    `gorm:"not null"`       // Quem executou a ação
    Type        string    `gorm:"not null"`       // "like", "comment"
    PostID      uint      `gorm:"index"`
    CommentID   *uint     `gorm:"index"`          // Nullable
    IsRead      bool      `gorm:"default:false"`
    CreatedAt   time.Time
    
    // Relacionamentos
    User    User  `gorm:"foreignKey:UserID"`
    Actor   User  `gorm:"foreignKey:ActorID"`
    Post    Post  `gorm:"foreignKey:PostID"`
}
```

---

## 🔌 API Endpoints (REST)

### Autenticação
Todos os endpoints (exceto públicos) precisam de header:
```
Authorization: Bearer <firebase-token>
```

### Posts
```
GET    /api/posts              - Lista todos posts (paginado)
GET    /api/posts/:id          - Detalhes de um post
POST   /api/posts              - Criar post
PUT    /api/posts/:id          - Editar post
DELETE /api/posts/:id          - Deletar post
GET    /api/users/:id/posts    - Posts de um usuário
```

### Likes
```
POST   /api/posts/:id/like     - Curtir post
DELETE /api/posts/:id/like     - Descurtir post
GET    /api/posts/:id/likes    - Lista quem curtiu
```

### Comments
```
GET    /api/posts/:id/comments     - Lista comentários
POST   /api/posts/:id/comments     - Criar comentário
PUT    /api/comments/:id           - Editar comentário
DELETE /api/comments/:id           - Deletar comentário
```

### Notifications
```
GET    /api/notifications          - Lista notificações do usuário
PUT    /api/notifications/:id/read - Marcar como lida
PUT    /api/notifications/read-all - Marcar todas como lidas
```

### Users
```
GET    /api/users/:id          - Perfil de usuário
PUT    /api/users/me           - Atualizar próprio perfil
```

---

## 🚀 Roadmap de Implementação

### Fase 1: Setup Inicial (Semana 1)
1. **Backend:**
   - [ ] Inicializar projeto Go (`go mod init`)
   - [ ] Instalar dependências (Gin, GORM, Firebase Admin SDK)
   - [ ] Configurar conexão PostgreSQL local
   - [ ] Criar middleware de autenticação Firebase
   - [ ] Setup CORS

2. **Frontend:**
   - [ ] Criar projeto Next.js (`npx create-next-app@latest`)
   - [ ] Configurar Tailwind CSS
   - [ ] Setup Firebase Authentication
   - [ ] Criar layout base

### Fase 2: Autenticação (Semana 2)
1. **Frontend:**
   - [ ] Página de login
   - [ ] Página de cadastro
   - [ ] Context de autenticação
   - [ ] Protected routes

2. **Backend:**
   - [ ] Middleware de validação de token
   - [ ] Endpoint de criação/atualização de user
   - [ ] Endpoint de profile

### Fase 3: Posts (Semana 3)
1. **Backend:**
   - [ ] Model Post + migrations
   - [ ] CRUD de posts
   - [ ] Validação de limite de caracteres
   - [ ] Paginação

2. **Frontend:**
   - [ ] Feed de posts
   - [ ] Formulário de criar post
   - [ ] Editar/deletar post
   - [ ] Validação de caracteres em tempo real

### Fase 4: Curtidas (Semana 4)
1. **Backend:**
   - [ ] Model Like + migrations
   - [ ] Endpoints de like/unlike
   - [ ] Constraint de unicidade
   - [ ] Criar notificação ao curtir

2. **Frontend:**
   - [ ] Botão de curtir
   - [ ] Contador de curtidas
   - [ ] Animação de curtida

### Fase 5: Comentários (Semana 5)
1. **Backend:**
   - [ ] Model Comment + migrations
   - [ ] CRUD de comentários
   - [ ] Criar notificação ao comentar

2. **Frontend:**
   - [ ] Lista de comentários
   - [ ] Formulário de comentar
   - [ ] Editar/deletar comentário

### Fase 6: Notificações (Semana 6)
1. **Backend:**
   - [ ] Model Notification + migrations
   - [ ] Endpoints de notificações
   - [ ] Lógica de marcar como lida

2. **Frontend:**
   - [ ] Badge de notificações não lidas
   - [ ] Lista de notificações
   - [ ] Polling de notificações

### Fase 7: Deploy (Semana 7)
1. **Banco de Dados:**
   - [ ] Criar projeto no Supabase
   - [ ] Rodar migrations no Supabase
   - [ ] Copiar DATABASE_URL

2. **Backend:**
   - [ ] Criar projeto no Railway
   - [ ] Configurar variáveis de ambiente
   - [ ] Configurar CORS para domínio Vercel
   - [ ] Deploy

3. **Frontend:**
   - [ ] Push para GitHub
   - [ ] Conectar com Vercel
   - [ ] Configurar variáveis de ambiente (API URL)
   - [ ] Deploy

### Fase 8: Refinamentos (Semana 8)
- [ ] Melhorias de UI/UX
- [ ] Loading states
- [ ] Error handling
- [ ] Testes básicos
- [ ] SEO básico

---

## 🔐 Variáveis de Ambiente

### Backend (Railway)
```env
DATABASE_URL=postgresql://user:password@host:5432/dbname
FIREBASE_PROJECT_ID=seu-projeto
FIREBASE_CREDENTIALS=<json-do-service-account>
PORT=8080
ENVIRONMENT=production
ALLOWED_ORIGINS=https://seu-app.vercel.app
```

### Frontend (Vercel)
```env
NEXT_PUBLIC_API_URL=https://seu-backend.railway.app
NEXT_PUBLIC_FIREBASE_API_KEY=...
NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN=...
NEXT_PUBLIC_FIREBASE_PROJECT_ID=...
NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET=...
NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID=...
NEXT_PUBLIC_FIREBASE_APP_ID=...
```

---

## 🎨 Sugestões de UI/UX

1. **Feed:**
   - Cards de posts com sombra suave
   - Avatar circular do autor
   - Timestamp relativo ("há 2 horas")
   - Botões de ação discretos

2. **Criar Post:**
   - Textarea com contador de caracteres
   - Mostrar limite visualmente (muda de cor perto do limite)
   - Botão desabilitado se vazio ou exceder limite

3. **Notificações:**
   - Badge vermelho com número
   - Dropdown ou página separada
   - Destacar não lidas

4. **Tema:**
   - Considerar dark mode
   - Cores: Azul/Roxo para ações, Vermelho para likes

---

## ⚠️ Pontos de Atenção

### Segurança
- [ ] Sempre validar token Firebase no backend
- [ ] Validar que usuário só pode editar/deletar próprio conteúdo
- [ ] Sanitizar inputs (prevenir XSS)
- [ ] Rate limiting (prevenir spam)

### Performance
- [ ] Índices no banco (UserID, PostID, CreatedAt)
- [ ] Paginação em todas as listas
- [ ] Cache de posts populares (opcional)
- [ ] Lazy loading de imagens de avatar

### Custos (Planos Gratuitos)
- **Vercel:** Grátis ilimitado para projetos pessoais ✅
- **Railway:** $5 de crédito/mês (~500 horas) ⚠️
- **Supabase:** 500MB, 2GB transferência ✅
- **Firebase Auth:** 50k MAU grátis ✅

**⚠️ IMPORTANTE:** Se o backend ficar 24/7 no Railway, consumirá ~730 horas/mês (excede grátis). Railway cobra $0.01/hora extra.

**Solução:** Configurar sleep mode ou migrar para Render.com (grátis mas mais lento).

---

## 📚 Recursos de Aprendizado

### Deploy
- Railway Docs: https://docs.railway.app/
- Supabase Quickstart: https://supabase.com/docs/guides/database
- Vercel Next.js Deploy: https://vercel.com/docs

### Go + GORM
- GORM Docs: https://gorm.io/docs/
- Gin Framework: https://gin-gonic.com/docs/

### Next.js
- Next.js 14: https://nextjs.org/docs
- Firebase Auth + Next.js: https://firebase.google.com/docs/auth/web/start

---

## 🤔 Decisões em Aberto

1. **WebSocket vs Polling para notificações?**
   - Recomendação: Começar com polling (mais simples)
   
2. **Comentários aninhados (respostas a comentários)?**
   - Recomendação: Deixar para v2 (aumenta complexidade)

3. **Upload de imagens nos posts?**
   - Não mencionado, assumindo apenas texto
   - Se quiser: Firebase Storage

4. **Sistema de seguidores?**
   - Não mencionado, mas seria útil
   - Deixar para v2?

5. **Railway vs Render para backend?**
   - Railway: Mais fácil, mas pago após limite
   - Render: Grátis, mas cold start (dorme após inatividade)

---

## 📝 Notas Finais

**O que tenho 90%+ certeza:**
- ✅ Stack escolhida funciona bem juntos
- ✅ Firebase + Go funciona (já usei)
- ✅ Next.js + Vercel é trivial
- ✅ GORM + PostgreSQL é maduro e estável

**O que tenho 70-85% certeza:**
- ⚠️ Railway para Go (funciona, mas nunca usei pessoalmente)
- ⚠️ Supabase como DB separado (funciona, mas conexão externa pode ter latência)
- ⚠️ Custos de Railway no longo prazo

**Recomendação:** Comece simples, implemente funcionalidades core primeiro, depois adicione complexidade.

**Próximos Passos Sugeridos:**
1. Revisar este plano juntos
2. Criar repositório GitHub
3. Setup do ambiente local
4. Começar pela Fase 1 do roadmap

---

**Dúvidas ou quer ajustar algo? Me avisa! 🚀**

