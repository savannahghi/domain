package invites

// TODO Gap: end to end invite process

// 1. Review: invite client by healthcare worker
// 2. Review: indicate that patient is counselled when HCW is inviting
// 3. Review: accept terms and conditions
// 4. Review: ensure that we can force a PIN change e.g for first time PIN
// 5. Review: PIN change API
// 6. Review: PIN hashing
// 7. Review: invite links: deep, tied to user details, client CCC, sent via SMS, expire
// 8. Review: send SMS API
// 9. Review: accept/reject invite
// 10. Review: set/update security questions
// 11. Review: verify security questions
// 12. Review: push token registration
// 13. Review: ability to choose onboarding questionnaire by length of treatment
// 14. Review: re-send invites (?invalidate first invite)
// 15. Review: send and resend OTP (different channel for resends)
// 16. Review: reset PIN, notify after PIN reset (both client reset and HCW reset)
// 17. Review: collecting metrics for key actions
// 18. Review: if a PIN is set by HCW or invite, it must be changed at next login
// 19. Review: ability to hook in additional rules on PIN reset e.g max resets
