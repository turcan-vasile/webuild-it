(() => {
  "use strict";

  const translations = {
    en: {
      "meta.title": "WeBuildit — AI, Automation & Digital Systems",
      "meta.description": "WeBuildit builds practical AI-assisted solutions, automation, digital products and reliable infrastructure for real businesses and ideas.",
      "meta.imageAlt": "WeBuildit — We build useful things together.",
      "a11y.skip": "Skip to main content",
      "a11y.primaryNav": "Primary navigation",
      "a11y.language": "Choose language",
      "nav.build": "What we build", "nav.work": "How we work", "nav.why": "Why WeBuildit", "nav.contact": "Contact",
      "hero.eyebrow": "AI • Automation • Digital Systems", "hero.title": "We build useful things together.",
      "hero.lead": "AI-assisted development, automation and digital systems for real businesses, products and ideas.",
      "hero.primary": "Start a project", "hero.secondary": "Explore what we build",
      "hero.note": "Human direction. Practical technology. Evidence-led improvement.",
      "capabilities.eyebrow": "What we build", "capabilities.title": "Useful systems, built around real needs.",
      "capabilities.lead": "From the first workflow to the production foundation, we connect the parts that make digital work dependable.",
      "capabilities.ai.title": "AI & Automation", "capabilities.ai.body": "Turn repetitive work and fragmented workflows into practical AI-assisted systems.",
      "capabilities.products.title": "Digital Products", "capabilities.products.body": "Websites, applications, internal tools and integrations built around real needs.",
      "capabilities.devops.title": "Infrastructure & DevOps", "capabilities.devops.body": "Reliable deployment, containers, CI/CD, GitOps, cloud and production operations.",
      "capabilities.security.title": "Security & Operations", "capabilities.security.body": "Hardening, monitoring, backups and safer operational foundations.",
      "capabilities.systems.title": "Business Systems", "capabilities.systems.body": "Connect tools, data and processes into systems that make everyday work simpler.",
      "process.eyebrow": "How we work", "process.title": "Understand → Build → Validate → Improve",
      "process.lead": "Delivery is a learning loop, not a handoff into the unknown.",
      "process.understand.title": "Understand", "process.understand.body": "Start with the real problem, constraints and desired outcome.",
      "process.build.title": "Build", "process.build.body": "Combine human expertise, software engineering and AI-assisted workflows.",
      "process.validate.title": "Validate", "process.validate.body": "Test the result against real use, reliability and security requirements.",
      "process.improve.title": "Improve", "process.improve.body": "Learn from evidence and iterate rather than treating delivery as the end.",
      "why.eyebrow": "Why WeBuildit", "why.title": "WeBuildit means we build it together.",
      "why.body1": "“We” can include the client who knows the problem, specialists who bring the right craft, and technology that expands what a small team can do.",
      "why.body2": "AI works inside a human-directed system. It accelerates research, implementation and review; people keep responsibility, context and judgment.",
      "why.principleLabel": "Our principle", "why.principle": "Proof before promise. Experience before hype.",
      "proof.eyebrow": "Real work", "proof.title": "Built in real systems.",
      "proof.lead": "This is the foundation for detailed case studies: clear context, honest constraints and verifiable outcomes.",
      "proof.ra.type": "Operating project", "proof.ra.body": "A real commerce and content ecosystem where web, automation and operations meet everyday business needs.",
      "proof.delivery.type": "Delivery systems", "proof.delivery.title": "Infrastructure & DevOps", "proof.delivery.body": "Container, CI/CD, GitOps and operational workflows developed against production constraints.",
      "proof.ai.type": "Working practice", "proof.ai.title": "AI-assisted engineering", "proof.ai.body": "Human-led workflows that use AI for analysis, implementation and review with explicit governance.",
      "community.eyebrow": "The next layer", "community.title": "Built by people. Extended by AI.",
      "community.body": "WeBuildit is evolving toward collaborative building with clients, trusted specialists and AI-supported workflows. The community is a direction we are developing—not a service we pretend already exists.",
      "contact.eyebrow": "Start a conversation", "contact.title": "Tell us what you want to build.",
      "contact.lead": "Share the problem, the idea or the system you want to improve. A short, honest description is enough to begin.",
      "contact.asideLabel": "What happens next", "contact.asideBody": "We read every request, check whether we are a useful fit and reply by email.",
      "form.name": "Name", "form.email": "Email", "form.company": "Company / project (optional)", "form.projectType": "Project type",
      "form.choose": "Choose a project type", "form.types.ai": "AI / Automation", "form.types.product": "Website / Digital Product",
      "form.types.devops": "Infrastructure / DevOps", "form.types.security": "Security / Operations", "form.types.other": "Other",
      "form.message": "Message", "form.messageHelp": "20–5,000 characters. Please do not include passwords or other secrets.",
      "form.submit": "Send project request", "form.sending": "Sending your request…",
      "form.success": "Thanks — your request has been received. We'll get back to you by email.",
      "form.validation": "Please check the highlighted fields and add a little more detail.",
      "form.rate": "Too many requests were sent from this connection. Please wait ten minutes and try again.",
      "form.unavailable": "The contact service is temporarily unavailable. Please try again later; your message has not been sent.",
      "form.error": "We could not deliver your request. Please try again; your message has not been sent.",
      "footer.line": "AI • Automation • Digital Systems • AI-assisted Development", "footer.rights": "All rights reserved."
    },
    ro: {
      "meta.title": "WeBuildit — AI, automatizare și sisteme digitale",
      "meta.description": "WeBuildit construiește soluții practice asistate de AI, automatizări, produse digitale și infrastructură fiabilă pentru afaceri și idei reale.",
      "meta.imageAlt": "WeBuildit — Construim împreună lucruri utile.",
      "a11y.skip": "Sari la conținutul principal", "a11y.primaryNav": "Navigare principală", "a11y.language": "Alege limba",
      "nav.build": "Ce construim", "nav.work": "Cum lucrăm", "nav.why": "De ce WeBuildit", "nav.contact": "Contact",
      "hero.eyebrow": "AI • Automatizare • Sisteme digitale", "hero.title": "Construim împreună lucruri utile.",
      "hero.lead": "Dezvoltare asistată de AI, automatizare și sisteme digitale pentru afaceri, produse și idei reale.",
      "hero.primary": "Începe un proiect", "hero.secondary": "Descoperă ce construim", "hero.note": "Direcție umană. Tehnologie practică. Îmbunătățire bazată pe dovezi.",
      "capabilities.eyebrow": "Ce construim", "capabilities.title": "Sisteme utile, create în jurul nevoilor reale.",
      "capabilities.lead": "De la primul flux de lucru până la fundația de producție, conectăm componentele care fac munca digitală fiabilă.",
      "capabilities.ai.title": "AI și automatizare", "capabilities.ai.body": "Transformăm munca repetitivă și fluxurile fragmentate în sisteme practice asistate de AI.",
      "capabilities.products.title": "Produse digitale", "capabilities.products.body": "Site-uri, aplicații, instrumente interne și integrări construite pentru nevoi reale.",
      "capabilities.devops.title": "Infrastructură și DevOps", "capabilities.devops.body": "Implementări fiabile, containere, CI/CD, GitOps, cloud și operațiuni de producție.",
      "capabilities.security.title": "Securitate și operațiuni", "capabilities.security.body": "Consolidare, monitorizare, copii de siguranță și fundații operaționale mai sigure.",
      "capabilities.systems.title": "Sisteme de business", "capabilities.systems.body": "Conectăm instrumente, date și procese în sisteme care simplifică munca de zi cu zi.",
      "process.eyebrow": "Cum lucrăm", "process.title": "Înțelegem → Construim → Validăm → Îmbunătățim", "process.lead": "Livrarea este o buclă de învățare, nu un transfer către necunoscut.",
      "process.understand.title": "Înțelegem", "process.understand.body": "Începem cu problema reală, constrângerile și rezultatul dorit.",
      "process.build.title": "Construim", "process.build.body": "Combinăm experiența umană, ingineria software și fluxurile de lucru asistate de AI.",
      "process.validate.title": "Validăm", "process.validate.body": "Testăm rezultatul în raport cu utilizarea reală, fiabilitatea și cerințele de securitate.",
      "process.improve.title": "Îmbunătățim", "process.improve.body": "Învățăm din dovezi și iterăm, în loc să tratăm livrarea ca pe un final.",
      "why.eyebrow": "De ce WeBuildit", "why.title": "WeBuildit înseamnă că îl construim împreună.",
      "why.body1": "„Noi” poate include clientul care cunoaște problema, specialiștii care aduc competențele potrivite și tehnologia care extinde posibilitățile unei echipe mici.",
      "why.body2": "AI funcționează într-un sistem condus de oameni. Accelerează cercetarea, implementarea și revizuirea; oamenii păstrează responsabilitatea, contextul și discernământul.",
      "why.principleLabel": "Principiul nostru", "why.principle": "Dovezi înaintea promisiunilor. Experiență înaintea exagerărilor.",
      "proof.eyebrow": "Muncă reală", "proof.title": "Construit în sisteme reale.", "proof.lead": "Aceasta este fundația viitoarelor studii de caz: context clar, constrângeri sincere și rezultate verificabile.",
      "proof.ra.type": "Proiect activ", "proof.ra.body": "Un ecosistem real de comerț și conținut, unde web-ul, automatizarea și operațiunile răspund nevoilor zilnice ale afacerii.",
      "proof.delivery.type": "Sisteme de livrare", "proof.delivery.title": "Infrastructură și DevOps", "proof.delivery.body": "Fluxuri de containere, CI/CD, GitOps și operațiuni dezvoltate în condiții reale de producție.",
      "proof.ai.type": "Practică de lucru", "proof.ai.title": "Inginerie asistată de AI", "proof.ai.body": "Fluxuri conduse de oameni care folosesc AI pentru analiză, implementare și revizuire, cu guvernanță explicită.",
      "community.eyebrow": "Următorul nivel", "community.title": "Construit de oameni. Extins prin AI.",
      "community.body": "WeBuildit evoluează spre construcție colaborativă alături de clienți, specialiști de încredere și fluxuri susținute de AI. Comunitatea este o direcție pe care o dezvoltăm, nu un serviciu despre care pretindem că există deja.",
      "contact.eyebrow": "Începe o conversație", "contact.title": "Spune-ne ce vrei să construiești.",
      "contact.lead": "Descrie problema, ideea sau sistemul pe care vrei să-l îmbunătățești. O explicație scurtă și sinceră este suficientă pentru început.",
      "contact.asideLabel": "Ce urmează", "contact.asideBody": "Citim fiecare solicitare, verificăm dacă putem fi de folos și răspundem prin e-mail.",
      "form.name": "Nume", "form.email": "E-mail", "form.company": "Companie / proiect (opțional)", "form.projectType": "Tipul proiectului",
      "form.choose": "Alege tipul proiectului", "form.types.ai": "AI / Automatizare", "form.types.product": "Site / Produs digital", "form.types.devops": "Infrastructură / DevOps", "form.types.security": "Securitate / Operațiuni", "form.types.other": "Altul",
      "form.message": "Mesaj", "form.messageHelp": "20–5.000 de caractere. Nu include parole sau alte date secrete.", "form.submit": "Trimite solicitarea", "form.sending": "Trimitem solicitarea…",
      "form.success": "Mulțumim — solicitarea ta a fost primită. Îți vom răspunde prin e-mail.", "form.validation": "Verifică câmpurile marcate și adaugă puțin mai multe detalii.",
      "form.rate": "Au fost trimise prea multe solicitări de la această conexiune. Așteaptă zece minute și încearcă din nou.",
      "form.unavailable": "Serviciul de contact este temporar indisponibil. Încearcă din nou mai târziu; mesajul nu a fost trimis.",
      "form.error": "Nu am putut livra solicitarea. Încearcă din nou; mesajul nu a fost trimis.",
      "footer.line": "AI • Automatizare • Sisteme digitale • Dezvoltare asistată de AI", "footer.rights": "Toate drepturile rezervate."
    },
    ru: {
      "meta.title": "WeBuildit — ИИ, автоматизация и цифровые системы",
      "meta.description": "WeBuildit создаёт практичные решения с применением ИИ, автоматизацию, цифровые продукты и надёжную инфраструктуру для реального бизнеса и идей.",
      "meta.imageAlt": "WeBuildit — Мы вместе создаём полезные вещи.",
      "a11y.skip": "Перейти к основному содержанию", "a11y.primaryNav": "Основная навигация", "a11y.language": "Выбрать язык",
      "nav.build": "Что мы создаём", "nav.work": "Как мы работаем", "nav.why": "Почему WeBuildit", "nav.contact": "Контакты",
      "hero.eyebrow": "ИИ • Автоматизация • Цифровые системы", "hero.title": "Мы вместе создаём полезные вещи.",
      "hero.lead": "Разработка с поддержкой ИИ, автоматизация и цифровые системы для реального бизнеса, продуктов и идей.",
      "hero.primary": "Начать проект", "hero.secondary": "Узнать, что мы создаём", "hero.note": "Человеческое руководство. Практичные технологии. Улучшения на основе фактов.",
      "capabilities.eyebrow": "Что мы создаём", "capabilities.title": "Полезные системы для реальных потребностей.",
      "capabilities.lead": "От первого рабочего процесса до производственной основы — мы соединяем компоненты, которые делают цифровую работу надёжной.",
      "capabilities.ai.title": "ИИ и автоматизация", "capabilities.ai.body": "Превращаем повторяющиеся задачи и разрозненные процессы в практичные системы с поддержкой ИИ.",
      "capabilities.products.title": "Цифровые продукты", "capabilities.products.body": "Сайты, приложения, внутренние инструменты и интеграции, созданные под реальные задачи.",
      "capabilities.devops.title": "Инфраструктура и DevOps", "capabilities.devops.body": "Надёжное развёртывание, контейнеры, CI/CD, GitOps, облачные и производственные операции.",
      "capabilities.security.title": "Безопасность и операции", "capabilities.security.body": "Защита, мониторинг, резервное копирование и более безопасная операционная основа.",
      "capabilities.systems.title": "Бизнес-системы", "capabilities.systems.body": "Объединяем инструменты, данные и процессы в системы, упрощающие повседневную работу.",
      "process.eyebrow": "Как мы работаем", "process.title": "Понимаем → Создаём → Проверяем → Улучшаем", "process.lead": "Поставка — это цикл обучения, а не передача результата в неизвестность.",
      "process.understand.title": "Понимаем", "process.understand.body": "Начинаем с реальной проблемы, ограничений и желаемого результата.",
      "process.build.title": "Создаём", "process.build.body": "Объединяем человеческий опыт, программную инженерию и рабочие процессы с поддержкой ИИ.",
      "process.validate.title": "Проверяем", "process.validate.body": "Проверяем результат в реальном использовании, а также требования к надёжности и безопасности.",
      "process.improve.title": "Улучшаем", "process.improve.body": "Учимся на фактах и дорабатываем решение, не считая поставку финальной точкой.",
      "why.eyebrow": "Почему WeBuildit", "why.title": "WeBuildit означает, что мы создаём это вместе.",
      "why.body1": "«Мы» — это клиент, который знает проблему, специалисты с подходящими навыками и технологии, расширяющие возможности небольшой команды.",
      "why.body2": "ИИ работает внутри системы, которой управляют люди. Он ускоряет исследования, реализацию и проверку; ответственность, контекст и решения остаются за людьми.",
      "why.principleLabel": "Наш принцип", "why.principle": "Сначала доказательства, потом обещания. Сначала опыт, потом громкие слова.",
      "proof.eyebrow": "Реальная работа", "proof.title": "Создано в реальных системах.", "proof.lead": "Это основа для подробных кейсов: ясный контекст, честные ограничения и проверяемые результаты.",
      "proof.ra.type": "Действующий проект", "proof.ra.body": "Реальная экосистема торговли и контента, где веб, автоматизация и операции решают повседневные задачи бизнеса.",
      "proof.delivery.type": "Системы поставки", "proof.delivery.title": "Инфраструктура и DevOps", "proof.delivery.body": "Контейнерные, CI/CD, GitOps и операционные процессы, разработанные с учётом производственных ограничений.",
      "proof.ai.type": "Рабочая практика", "proof.ai.title": "Инженерия с поддержкой ИИ", "proof.ai.body": "Управляемые людьми процессы, использующие ИИ для анализа, реализации и проверки в рамках явных правил.",
      "community.eyebrow": "Следующий уровень", "community.title": "Создано людьми. Расширено с помощью ИИ.",
      "community.body": "WeBuildit развивается в сторону совместной работы с клиентами, проверенными специалистами и процессами с поддержкой ИИ. Сообщество — это направление, которое мы развиваем, а не услуга, которая якобы уже существует.",
      "contact.eyebrow": "Начать разговор", "contact.title": "Расскажите, что вы хотите создать.",
      "contact.lead": "Опишите проблему, идею или систему, которую хотите улучшить. Для начала достаточно короткого и честного описания.",
      "contact.asideLabel": "Что будет дальше", "contact.asideBody": "Мы читаем каждую заявку, оцениваем, можем ли быть полезны, и отвечаем по электронной почте.",
      "form.name": "Имя", "form.email": "Электронная почта", "form.company": "Компания / проект (необязательно)", "form.projectType": "Тип проекта",
      "form.choose": "Выберите тип проекта", "form.types.ai": "ИИ / Автоматизация", "form.types.product": "Сайт / Цифровой продукт", "form.types.devops": "Инфраструктура / DevOps", "form.types.security": "Безопасность / Операции", "form.types.other": "Другое",
      "form.message": "Сообщение", "form.messageHelp": "От 20 до 5 000 символов. Не указывайте пароли и другие секретные данные.", "form.submit": "Отправить заявку", "form.sending": "Отправляем заявку…",
      "form.success": "Спасибо — ваша заявка получена. Мы ответим по электронной почте.", "form.validation": "Проверьте отмеченные поля и добавьте немного больше информации.",
      "form.rate": "С этого подключения отправлено слишком много заявок. Подождите десять минут и повторите попытку.",
      "form.unavailable": "Сервис связи временно недоступен. Повторите попытку позже; сообщение не было отправлено.",
      "form.error": "Не удалось доставить заявку. Повторите попытку; сообщение не было отправлено.",
      "footer.line": "ИИ • Автоматизация • Цифровые системы • Разработка с поддержкой ИИ", "footer.rights": "Все права защищены."
    }
  };

  const languages = ["en", "ro", "ru"];
  let currentLanguage = "en";

  const readStoredLanguage = () => {
    try {
      const stored = window.localStorage.getItem("webuildit-language");
      return languages.includes(stored) ? stored : "en";
    } catch (_error) {
      return "en";
    }
  };

  const updateMeta = (language) => {
    const dictionary = translations[language];
    document.title = dictionary["meta.title"];
    document.querySelector('meta[name="description"]').setAttribute("content", dictionary["meta.description"]);
    document.querySelector('meta[property="og:title"]').setAttribute("content", dictionary["meta.title"]);
    document.querySelector('meta[property="og:description"]').setAttribute("content", dictionary["meta.description"]);
    document.querySelector('meta[property="og:image:alt"]').setAttribute("content", dictionary["meta.imageAlt"]);
    document.querySelector('meta[name="twitter:title"]').setAttribute("content", dictionary["meta.title"]);
    document.querySelector('meta[name="twitter:description"]').setAttribute("content", dictionary["meta.description"]);
    document.querySelector('meta[name="twitter:image:alt"]').setAttribute("content", dictionary["meta.imageAlt"]);
  };

  const setLanguage = (language, persist = true) => {
    if (!languages.includes(language)) return;
    currentLanguage = language;
    document.documentElement.lang = language;

    document.querySelectorAll("[data-i18n]").forEach((element) => {
      const value = translations[language][element.dataset.i18n];
      if (value !== undefined) element.textContent = value;
    });
    document.querySelectorAll("[data-i18n-aria-label]").forEach((element) => {
      const value = translations[language][element.dataset.i18nAriaLabel];
      if (value !== undefined) element.setAttribute("aria-label", value);
    });
    document.querySelectorAll("[data-language]").forEach((button) => {
      const active = button.dataset.language === language;
      button.classList.toggle("is-active", active);
      button.setAttribute("aria-pressed", String(active));
    });
    updateMeta(language);

    if (persist) {
      try {
        window.localStorage.setItem("webuildit-language", language);
      } catch (_error) {
        // Language switching still works when storage is unavailable.
      }
    }
  };

  document.querySelectorAll("[data-language]").forEach((button) => {
    button.addEventListener("click", () => setLanguage(button.dataset.language));
  });

  const form = document.querySelector("#contact-form");
  const status = document.querySelector("#form-status");

  const setStatus = (key, type) => {
    status.textContent = translations[currentLanguage][key];
    status.className = `form-status ${type ? `is-${type}` : ""}`.trim();
  };

  form.addEventListener("submit", async (event) => {
    event.preventDefault();
    status.textContent = "";
    status.className = "form-status";

    if (!form.checkValidity()) {
      form.reportValidity();
      setStatus("form.validation", "error");
      return;
    }

    const submitButton = form.querySelector('button[type="submit"]');
    const submitLabel = submitButton.querySelector("[data-i18n]");
    submitButton.disabled = true;
    submitLabel.textContent = translations[currentLanguage]["form.sending"];

    const data = new FormData(form);
    const payload = Object.fromEntries(data.entries());
    payload.language = currentLanguage;

    try {
      const response = await fetch(form.action, {
        method: "POST",
        headers: { "Content-Type": "application/json", "Accept": "application/json" },
        body: JSON.stringify(payload)
      });

      if (response.ok) {
        form.reset();
        setStatus("form.success", "success");
      } else if (response.status === 400 || response.status === 413) {
        setStatus("form.validation", "error");
      } else if (response.status === 429) {
        setStatus("form.rate", "error");
      } else if (response.status === 503) {
        setStatus("form.unavailable", "error");
      } else {
        setStatus("form.error", "error");
      }
    } catch (_error) {
      setStatus("form.error", "error");
    } finally {
      submitButton.disabled = false;
      submitLabel.textContent = translations[currentLanguage]["form.submit"];
    }
  });

  document.querySelector("#current-year").textContent = new Date().getFullYear();
  setLanguage(readStoredLanguage(), false);
})();
