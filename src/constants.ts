export const CONTACT = Object.freeze({
	maxEmailLength: Number(process.env.CONTACT_MAX_EMAIL_ADDRESS_LENGTH) || 254,
	maxNameLength: Number(process.env.CONTACT_MAX_NAME_LENGTH) || 70,
	maxMessageLength: Number(process.env.CONTACT_MAX_MESSAGE_LENGTH) || 10000
});
