;;; package--- Summary
(require 'package) 
(add-to-list 'package-archives '("melpa" . "https://melpa.org/packages/")) 
(package-initialize)
;;; Commentary:xx
(setq inhibit-startup-message t)        ;close start page
(desktop-save-mode t)                   ;保存session
(electric-pair-mode t)                  ;括号补全
(custom-set-variables
 ;; custom-set-variables was added by Custom.
 ;; If you edit it by hand, you could mess it up, so be careful.
 ;; Your init file should contain only one such instance.
 ;; If there is more than one, they won't work right.
 '(global-linum-mode t))
(add-to-list 'load-path "~/.emacs.d/elpa")
(require 'elisp-format)
;;;vim mode
(require 'evil)
(evil-mode 1)
                                        ;neotree
(require 'neotree)
(global-set-key (kbd "<f5>") 'neotree-toggle)
(global-set-key (kbd "<f6>") 
                (lambda () 
                  (interactive) 
                  (neotree-dir "/var/www/html")))
(defun neotree-close-up-node () 
  (interactive) 
  (neotree-select-up-node) 
  (neotree-enter))

(defun neotree-current-to-root () 
  (interactive) 
  (neotree-change-root))

(defun neotree-open-vsplite () 
  (interactive) 
  (neotree-enter-vertical-split))

(defun neotree-open-splite () 
  (interactive) 
  (neotree-enter-horizontal-split))

(setq neo-smart-open t)
(add-hook 'neotree-mode-hook 
          (lambda () 
            (define-key evil-normal-state-local-map (kbd "TAB") 'neotree-enter) 
            (define-key evil-normal-state-local-map (kbd "o") 'neotree-enter) 
            (define-key evil-normal-state-local-map (kbd "i") 'neotree-open-splite) 
            (define-key evil-normal-state-local-map (kbd "s") 'neotree-open-vsplite) 
            (define-key evil-normal-state-local-map (kbd "t") 'neotree-current-to-root) 
            (define-key evil-normal-state-local-map (kbd "x") 'neotree-close-up-node) 
            (define-key evil-normal-state-local-map (kbd "SPC") 'neotree-enter) 
            (define-key evil-normal-state-local-map (kbd "q") 'neotree-hide) 
            (define-key evil-normal-state-local-map (kbd "RET") 'neotree-enter)))

(require 'highlight-parentheses)
(define-globalized-minor-mode global-highlight-parentheses-mode highlight-parentheses-mode 
  (lambda () 
    (highlight-parentheses-mode t)))
(global-highlight-parentheses-mode t)

(defun my-go-mode-hook ()
                                        ; Call Gofmt before saving
  (add-hook 'before-save-hook 'gofmt-before-save)
                                        ; Godef jump key binding
  (local-set-key (kbd "s-/") 'godef-jump) 
  (local-set-key (kbd "s-b") 'pop-tag-mark))
(add-hook 'go-mode-hook 'my-go-mode-hook)
(load-file "~/.emacs.d/.entry.el")
