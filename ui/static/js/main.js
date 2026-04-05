const body = document.body;
const revealItems = document.querySelectorAll("[data-reveal]");

const syncScrollState = () => {
  body.classList.toggle("is-scrolled", window.scrollY > 12);
};

syncScrollState();
window.addEventListener("scroll", syncScrollState, { passive: true });

if ("IntersectionObserver" in window) {
  const observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (!entry.isIntersecting) {
          return;
        }

        entry.target.classList.add("is-visible");
        observer.unobserve(entry.target);
      });
    },
    {
      rootMargin: "0px 0px -10% 0px",
      threshold: 0.12,
    },
  );

  revealItems.forEach((item) => observer.observe(item));
} else {
  revealItems.forEach((item) => item.classList.add("is-visible"));
}
